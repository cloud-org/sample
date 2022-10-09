package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id" bson:"id"`
	Name     string    `json:"name" bson:"name"`
	LastSeen time.Time `json:"lastSeen" bson:"lastSeen"`
}

func main() {
	a := &MyUser{1, "Ken", time.Now()}
	b, err := bson.Marshal(a)
	if err != nil{
		log.Println(err)
		return
	}
	log.Printf("b is %+v\n", bson.Raw(b))
	var c MyUser
	log.Println(bson.Unmarshal(b, &c))
	log.Printf("c is %+v\n", c)
}

func (u *MyUser) MarshalBSON() ([]byte, error) {
	type Alias MyUser
	return bson.Marshal(&struct {
		LastSeen int64 `json:"lastSeen" bson:"lastSeen"`
		*Alias
	}{
		LastSeen: u.LastSeen.Unix(),
		Alias:    (*Alias)(u),
	})
}

func (u *MyUser) UnmarshalBSON(data []byte) error {
	type Alias MyUser
	aux := &struct {
		LastSeen int64 `json:"lastSeen" bson:"lastSeen"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := bson.Unmarshal(data, &aux); err != nil {
		return err
	}
	log.Printf("aux is %+v\n", aux)
	log.Printf("aux is %+v\n", aux.Alias)
	u.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}
