// author: ashing
// time: 2020/4/26 11:10 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"context"
	"log"

	proto "github.com/ronething/grpc-sample/protoc3/protoc"
	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
