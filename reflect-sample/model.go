/*
 * MIT License
 *
 * Copyright (c) 2021 ashing
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package reflect_sample

import (
	"context"
	"errors"
	"log"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Validator interface {
	ValidateStruct() error
}

type ModelMethod struct {
	Col       *mongo.Collection
	ColName   string
	Request   *gin.Context
	MetaModel interface{}
	// TODO: 放在一个通用的位置
	v *validator.Validate
}

func NewModelMethod(col *mongo.Collection, colName string, metaModel interface{}, ctx *gin.Context) ModelMethod {
	return ModelMethod{
		Col:       col,
		ColName:   colName,
		MetaModel: metaModel,
		Request:   ctx,
		v:         validator.New(),
	}
}

func (m *ModelMethod) ValidateStruct(data interface{}) error {
	value := reflect.ValueOf(m.MetaModel)
	method := value.MethodByName("ValidateStruct") // 是否有对应的自定义 func
	if method.IsValid() {
		dataValue := reflect.ValueOf(data)
		args := []reflect.Value{dataValue}
		res := method.Call(args)
		err, ok := res[0].Interface().(error)
		if ok {
			return err
		} else {
			return errors.New("实现 func 有问题")
		}
	} else {
		return m.v.Struct(data)
	}
}

func optData(data map[string]interface{}, key string, value interface{}) {
	_, ok := data[key]
	if !ok || data[key] == "" || data[key] == nil {
		// 因为一般是设置时间或者其他的需要有值的字段 所以如果键不存在或者没有值 就设置一下
		data[key] = value
	}
}

type User struct {
	UserName string
}

func (m *ModelMethod) FillUserField(data map[string]interface{}) {
	if m.Request != nil {
		user, ok := m.Request.Get("user")
		if ok {
			currentUser, ok := user.(*User)
			if ok {
				optData(data, "creator", currentUser.UserName)
				optData(data, "last_changer", currentUser.UserName)
				return
			}
		}

		optData(data, "creator", "api")
		optData(data, "last_changer", "api")
		return
	}
	return
}

func (m *ModelMethod) FillTimeField(data map[string]interface{}) {
	optData(data, "ctime", time.Now())
	optData(data, "mtime", time.Now())
}

func (m *ModelMethod) Interface2json(inter interface{}, tagName string) map[string]interface{} {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("interface2json err: %v\n", err)
		}
		return
	}()
	data := make(map[string]interface{})
	t := reflect.TypeOf(inter)
	if t.Kind() == reflect.Ptr { // non-struct num field https://www.codenong.com/cs109293647/
		t = t.Elem()
		log.Printf("t kind is %v\n", t.Kind())
	}
	v := reflect.ValueOf(inter)
	//v = reflect.Indirect(v) // 等价于下面的
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		log.Printf("value is %v\n", v.Kind())
	}

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)
		tag = strings.Split(tag, ",")[0]
		if tag != "" { // 如果没有对应的 tag，则直接略过
			data[tag] = v.Field(i).Interface()
		}
	}

	return data
}

func (m *ModelMethod) ValidateAnd2json(data interface{}) (map[string]interface{}, error) {
	err := m.ValidateStruct(data)
	if err != nil {
		return nil, err
	}
	// TODO: 暂定 json tag
	return m.Interface2json(data, "json"), nil
}

func (m *ModelMethod) InsertOne(ctx context.Context, inter interface{}) (interface{}, error) {
	// 先校验
	data, err := m.ValidateAnd2json(inter)
	if err != nil {
		return nil, err
	}
	// 增加数据
	m.FillTimeField(data)
	m.FillUserField(data)

	res, err := m.Col.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (m *ModelMethod) UpdateOne(ctx context.Context, id string, data map[string]interface{}) (*mongo.UpdateResult, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	optData(data, "mtime", time.Now())
	res, err := m.Col.UpdateOne(ctx, _id, bson.M{
		"$set": data,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *ModelMethod) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (interface{}, error) {
	t := reflect.TypeOf(m.MetaModel)
	value := reflect.New(t).Elem().Interface() // Elem 值类型
	err := m.Col.FindOne(ctx, filter, opts...).Decode(&value)
	return value, err
}
