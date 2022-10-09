// author: ashing
// time: 2020/4/26 2:37 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"io"
	"log"
	"net"

	"github.com/ronething/grpc-sample/protoc4/proto"
	"google.golang.org/grpc"
)

type StreamService struct {
}

func (s StreamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s StreamService) Record(stream proto.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.StreamResponse{Pt: &proto.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	//return nil
}

func (s StreamService) Route(stream proto.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	//return nil
}

const Port = "9002"

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", "127.0.0.1:"+Port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
