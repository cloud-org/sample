// author: ashing
// time: 2020/4/6 10:00 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/ronething/grpc-sample/constant"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call(constant.HelloServiceName+".Hello", "panda", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
