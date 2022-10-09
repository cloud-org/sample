package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// 真实需求场景
type LimitCondItem struct {
	Field     string      `json:"field" bson:"field"`
	Operator  string      `json:"operator" bson:"operator"`
	Value     interface{} `json:"value" bson:"value"` // 查询结果反序列化的时候不要使用 有坑 {Key: xx, Value: xx}
	FieldType string      `json:"fieldType" bson:"fieldType"`
	ValueType string      `json:"valueType" bson:"valueType"`
}

type GetOperator struct {
	Operator string `json:"operator" bson:"operator"`
}

type Range struct {
	Start int `json:"start" bson:"start"`
	End   int `json:"end" bson:"end"`
}

func (l *LimitCondItem) UnmarshalBSON(data []byte) (err error) {
	log.Println(bson.Raw(data))
	var g GetOperator
	if err = bson.Unmarshal(data, &g); err != nil {
		return err
	}
	type Alias LimitCondItem
	switch g.Operator {
	case "$range":
		aux := &struct {
			Value  Range `bson:"value"`
			*Alias `bson:",inline"` // 内联关键
		}{
			Alias: (*Alias)(l),
		}
		if err = bson.Unmarshal(data, &aux); err != nil {
			return err
		}
		l.Value = aux.Value // 赋值
		return
	case "$rangeArray":
		aux := &struct {
			Value  []Range `bson:"value"`
			*Alias `bson:",inline"`
		}{
			Alias: (*Alias)(l),
		}
		if err = bson.Unmarshal(data, &aux); err != nil {
			return err
		}
		l.Value = aux.Value
		return
	default:
		return bson.Unmarshal(data, (*Alias)(l))
	}
}
