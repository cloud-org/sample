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
	"log"
	"reflect"
	"testing"
)

type Hello struct {
	Name  string `json:"name" bson:"name" validate:"required"`
	Age   int    `json:"age" bson:"age" validate:"gt=0"`
	Ctime string `json:"ctime"`
	Mtime string `json:"mtime"`
}

func NewHello(name string, age int) *Hello {
	return &Hello{
		Name: name,
		Age:  age,
	}
}

// 自定义校验
//func (h *Hello) ValidateStruct(data interface{}) error {
//	log.Println("data is", data)
//	return errors.New("validate error")
//}

func TestModelMethod_ValidateStruct(t *testing.T) {
	var hello Hello
	m := NewModelMethod(nil, "hello", hello, nil)
	// 进行数据的补变动 然后进行结构体的校验
	//h.Age = 19
	//err := h.ValidateStruct(h)
	hello.Name = "panda"
	hello.Age = 0
	err := m.ValidateStruct(hello)
	if err != nil {
		t.Error(err)
	}
}

func TestModelMethod_Interface2json(t *testing.T) {
	//var hello Hello
	hello := NewHello("panda", 1)
	m := NewModelMethod(nil, "hello", hello, nil)
	// 进行数据的变动 然后进行结构体的校验
	data, err := m.ValidateAnd2json(hello)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("data is", data)
}

type HelloWorld struct {
	Name string `json:"name" bson:"name" validate:"required"`
	Age  int    `json:"age" bson:"age" validate:"gt=0"`
}

func TestMongoBsonM(t *testing.T) {
	result := buildResult()
	h, ok := result.(HelloWorld)
	if !ok {
		t.Error("err")
		return
	}
	t.Log("h is", h)
}

func buildResult() interface{} {
	h := HelloWorld{}
	t := reflect.TypeOf(h)
	value := reflect.New(t).Elem().Interface()
	log.Printf("value is %v\n", value)
	v := value.(HelloWorld)
	v.Age = 18
	v.Name = "panda"
	return v
}
