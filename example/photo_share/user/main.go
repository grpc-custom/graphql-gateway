package main

import (
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/photo_share/proto/user"
	"github.com/grpc-custom/graphql-gateway/example/photo_share/user/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.NewUserService()
	user.RegisterUserServerServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
