package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cloud-org/msgpush"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("listen 9000")
	http.ListenAndServe(":9000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	requestId := r.URL.Query().Get("requestId")
	log.Println("reqId is", requestId)
	go SendMsg(requestId)
	fmt.Fprintf(w, "Hello, Golang! from hook")
}

func SendMsg(id string) {
	mp := msgpush.NewPushDeer("")
	value := time.Now().String()
	err := mp.Send(id + " " + value)
	log.Println("err is", err)
}
