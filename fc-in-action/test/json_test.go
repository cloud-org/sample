package test

import (
	"encoding/json"
	"testing"
)

type Request struct {
	Event   string `json:"event"`
	Payload []byte `json:"payload"` // json 字符串
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonMarshal(t *testing.T) {
	stu := Student{
		Name: "panda",
		Age:  18,
	}
	data, err := json.Marshal(&stu)
	if err != nil {
		t.Error("marshal", err)
		return
	}
	a := Request{
		Event:   "order",
		Payload: data,
	}
	data, err = json.Marshal(&a)
	if err != nil {
		t.Error("marshal", err)
		return
	}

	t.Log("data is", string(data))

	var request Request
	err = json.Unmarshal(data, &request)
	if err != nil {
		t.Error("unmarshal", err)
		return
	}

	t.Logf("%+v\n", request.Event)

	var stud Student
	err = json.Unmarshal(request.Payload, &stud)
	if err != nil {
		t.Error("unmarshal", err)
		return
	}

	t.Logf("%+v\n", stud)

}
