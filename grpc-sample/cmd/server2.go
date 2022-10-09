// author: ashing
// time: 2020/4/6 10:04 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/ronething/grpc-sample/service2"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	time.Sleep(5 * time.Second) // 睡眠 5 秒
	*reply = "hello: " + request
	return nil
}

func main() {
	service2.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	log.Printf("server on port: 1234\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		//rpc.ServeConn(conn) 如果有多个 client 请求，会阻塞，所以使用 go func 启动 goroutine 进行处理
		go rpc.ServeConn(conn)
	}
}
