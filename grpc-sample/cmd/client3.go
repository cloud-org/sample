// author: ashing
// time: 2020/4/6 11:01 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "jsonrpc panda", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
