// author: ashing
// time: 2020/4/6 10:01 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"log"

	"github.com/ronething/grpc-sample/client2"
)

func main() {
	client, err := client2.DialHelloService("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("panda", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
