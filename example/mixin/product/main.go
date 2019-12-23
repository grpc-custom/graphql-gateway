package main

import (
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/mixin/product/service"
	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/product"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.New()
	product.RegisterProductServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
