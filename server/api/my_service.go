package api

import (
	"context"
	pb "grpc-gateway-101/server/proto"

	"google.golang.org/grpc/codes"
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

func (s *myServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {

	result := &pb.HealthCheckResponse{
		Code:    codes.OK.String(),
		Message: "API is running!",
	}

	return result, nil

}
