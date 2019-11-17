package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/photo_share/photo/service"
	"github.com/grpc-custom/graphql-gateway/example/photo_share/proto/photo"
	"google.golang.org/grpc"
)

func main() {
	var (
		port = flag.Int("port", 9001, "port number")
	)
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	svc := service.NewPhotoService()
	photo.RegisterPhotoServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
