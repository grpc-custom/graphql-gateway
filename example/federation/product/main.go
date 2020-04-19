package main

import (
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/federation/product/service"
	"github.com/grpc-custom/graphql-gateway/example/federation/proto/product"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.NewProductService()
	product.RegisterProductServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
