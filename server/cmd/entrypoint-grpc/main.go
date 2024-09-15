package main

import (
	"fmt"
	"grpc-gateway-101/server/api"
	"net"

	pb "grpc-gateway-101/server/proto"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

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

	grpcMuxServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	pb.RegisterMyServiceServer(
		grpcMuxServer, srv,
	)

	grpc_prometheus.Register(grpcMuxServer)

	fmt.Println("[main][func: RunGRPC] gRPC Server is running on port 9000")

	err = grpcMuxServer.Serve(listener)
	if err != nil {
		fmt.Println("[main][func: RunGRPC] Couldn't serve grpc listener")
		return
	}

}
