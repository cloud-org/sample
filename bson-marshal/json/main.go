package main

import (
	"encoding/json"
	"log"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	//_ = json.NewEncoder(os.Stdout).Encode(
	//	&MyUser{1, "Ken", time.Now()},
	//)
	a := &MyUser{1, "Ken", time.Now()}
	b, err := json.Marshal(a)
	if err != nil{
		log.Println(err)
		return
	}
	var c MyUser
	log.Println(json.Unmarshal(b, &c))
}

func (u *MyUser) MarshalJSON() ([]byte, error) {
	type Alias MyUser
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		LastSeen: u.LastSeen.Unix(),
		Alias:    (*Alias)(u),
	})
}

func (u *MyUser) UnmarshalJSON(data []byte) error {
	type Alias MyUser
	aux := &struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	log.Printf("aux is %+v\n", aux)
	log.Printf("aux is %+v\n", aux.Alias)
	u.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}
