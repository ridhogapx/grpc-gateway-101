package main

import (
	"context"
	"fmt"
	"net/http"

	pb "grpc-gateway-101/server/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	fmt.Println("[main][func: RunGateway] Starting gateway server on port :8080")

	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	client := pb.NewMyServiceClient(conn)

	ctx := context.Background()
	rmux := runtime.NewServeMux()

	err = pb.RegisterMyServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		fmt.Println("[main][func: RunGateway] Couldn't register service from endpoint:", err)
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", rmux)

	mdlw := middleware.New(middleware.Config{
		Recorder:           metrics.NewRecorder(metrics.Config{}),
		DisableMeasureSize: true,
	})

	h := std.Handler("", mdlw, mux)

	http.ListenAndServe(":8080", h)

}
