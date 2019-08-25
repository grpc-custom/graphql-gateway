package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/sample/proto/green"
	"github.com/grpc-custom/graphql-gateway/example/sample/proto/red"
	"github.com/grpc-custom/graphql-gateway/example/sample/service"
	"google.golang.org/grpc"
)

func main() {
	var (
		port = flag.Int("port", 9000, "port")
		svc  = flag.String("svc", "", "service name")
	)
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	switch *svc {
	case "red":
		registerRedServer(server)
	case "green":
		registerGreenServer(server)
	default:
		panic("unimplemented gRPC server")
	}
	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func registerRedServer(server *grpc.Server) {
	svc := service.NewReadService()
	red.RegisterRedServiceServer(server, svc)
}

func registerGreenServer(server *grpc.Server) {
	svc := service.NewGreenService()
	green.RegisterGreenServiceServer(server, svc)
}
