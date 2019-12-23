package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/account"
	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/product"
	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/review"
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
	err = review.RegisterReviewServiceFromEndpoint(ctx, mux, "localhost:9002", opts)
	if err != nil {
		log.Fatal(err)
	}
	err = account.RegisterAccountServiceFromEndpoint(ctx, mux, "localhost:9001", opts)
	if err != nil {
		log.Fatal(err)
	}
	err = product.RegisterProductServiceFromEndpoint(ctx, mux, "localhost:9003", opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
