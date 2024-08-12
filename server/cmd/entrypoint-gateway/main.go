package main

import (
	"context"
	"fmt"
	"net/http"

	pb "grpc-gateway-101/server/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	fmt.Println("[main][func: RunGateway] Starting gateway server on port :3000")

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx := context.Background()

	err := pb.RegisterMyServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts)
	if err != nil {
		fmt.Println("[main][func: RunGateway] Couldn't register service from endpoint:", err)
		return
	}

	http.ListenAndServe(":3000", mux)

}
