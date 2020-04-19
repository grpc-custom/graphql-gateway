package main

import (
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/federation/proto/review"
	"github.com/grpc-custom/graphql-gateway/example/federation/review/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.NewReviewService()
	review.RegisterReviewServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
