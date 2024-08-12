package main

import (
	"fmt"
	"grpc-gateway-101/server/api"
	"net"

	pb "grpc-gateway-101/server/proto"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("[main][func: RunGRPC] Starting gRPC Server...")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("[main][func: RunGRPC] Couldn't listen tcp in port 9000")
		return
	}

	srv := api.New()

	grpcServer := grpc.NewServer()
	pb.RegisterMyServiceServer(
		grpcServer, srv,
	)

	fmt.Println("[main][func: RunGRPC] gRPC Server is running on port 9000")

	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println("[main][func: RunGRPC] Couldn't serve grpc listener")
		return
	}

}
