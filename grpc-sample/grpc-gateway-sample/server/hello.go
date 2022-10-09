// author: ashing
// time: 2020/7/12 2:49 下午
// mail: axingfly@gmail.com
// Less is more.

package server

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/ronething/grpc-sample/grpc-gateway-sample/proto"
)

type helloService struct{}

func NewHelloService() *helloService {
	return &helloService{}
}

func (h helloService) SayHelloWorld(ctx context.Context, r *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: fmt.Sprintf("hello, %s", r.Referer),
	}, nil
}
