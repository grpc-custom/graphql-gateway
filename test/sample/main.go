package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-custom/graphql-gateway/runtime"
	"github.com/grpc-custom/graphql-gateway/test/sample/proto/red"
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
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
