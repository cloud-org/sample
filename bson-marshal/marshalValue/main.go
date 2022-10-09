package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type URL struct {
	URI string
}

func (u *URL) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(u.URI)
}

func (u *URL) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	log.Println(string(data))
	rv := bson.RawValue{Type: t, Value: data}
	return rv.Unmarshal(&u.URI)
}

type MyStruct struct {
	Image URL `bson:"image"`
}

func main() {
	m := &MyStruct{
		Image: URL{
			URI: "foobar",
		},
	}
	res, err := bson.Marshal(m)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n",bson.Raw(res))
	fmt.Printf("%+v\n",bson.Raw(res))

	var unmarshalled MyStruct
	if err := bson.Unmarshal(res, &unmarshalled); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", unmarshalled)
}