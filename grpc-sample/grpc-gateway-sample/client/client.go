// author: ashing
// time: 2020/7/12 3:01 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/ronething/grpc-sample/grpc-gateway-sample/proto"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../certs/server.pem", "localhost")
	if err != nil {
		log.Printf("Failed to create TLS credentials %v\n", err)
		return
	}
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))
	defer conn.Close()

	if err != nil {
		log.Printf("Dial err: %v\n", err)
		return
	}

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer: "Grpc",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(r.Message)
}
