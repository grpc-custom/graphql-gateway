package main

import (
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/federation/account/service"
	"github.com/grpc-custom/graphql-gateway/example/federation/proto/account"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.NewAccountService()
	account.RegisterAccountServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
