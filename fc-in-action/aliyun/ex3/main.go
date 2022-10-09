// +build linux

package main

import (
	"fmt"
	"log"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.Start(HandleRequest)
}

func HandleRequest(event string) (string, error) {
	log.Println("hello world from hook log")
	fmt.Println("hello world")
	fmt.Printf("event: %s\n", event)
	return "hello world", nil
}
