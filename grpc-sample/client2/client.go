// author: ashing
// time: 2020/4/6 10:14 下午
// mail: axingfly@gmail.com
// Less is more.

package client2

import (
	"net/rpc"

	"github.com/ronething/grpc-sample/constant"
	"github.com/ronething/grpc-sample/service2"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ service2.HelloServiceInterface = (*HelloServiceClient)(nil) // 确保 client 有实现 interface

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(constant.HelloServiceName+".Hello", request, reply)
}
