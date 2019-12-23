package main

import (
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/mixin/account/service"
	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/account"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.New()
	account.RegisterAccountServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
