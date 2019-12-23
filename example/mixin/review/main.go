package main

import (
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/review"
	"github.com/grpc-custom/graphql-gateway/example/mixin/review/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.New()
	review.RegisterReviewServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
