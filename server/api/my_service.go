package api

import (
	"context"
	pb "grpc-gateway-101/server/proto"
)

type myServer struct {
	pb.MyServiceServer
}

func New() *myServer {
	return &myServer{}
}

func (s myServer) Echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {

	result := &pb.StringMessage{
		Value: "Echoing....",
	}

	return result, nil

}
