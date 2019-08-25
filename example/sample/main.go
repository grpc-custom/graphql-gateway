package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-custom/graphql-gateway/example/sample/proto/green"
	"github.com/grpc-custom/graphql-gateway/example/sample/proto/red"
	"github.com/grpc-custom/graphql-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux, err := runtime.NewServeMux()
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	redHost := "localhost:9001"
	err = red.RegisterRedServiceFromEndpoint(ctx, mux, redHost, opts)
	if err != nil {
		log.Fatal(err)
	}
	greenHost := "localhost:9002"
	err = green.RegisterGreenServiceFromEndpoint(ctx, mux, greenHost, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
