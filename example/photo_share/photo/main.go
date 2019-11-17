package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-custom/graphql-gateway/example/photo_share/photo/service"
	"github.com/grpc-custom/graphql-gateway/example/photo_share/proto/photo"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			auth.UnaryServerInterceptor(authFunc),
		)),
	)
	svc := service.NewPhotoService()
	photo.RegisterPhotoServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func authFunc(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	return ctx, nil
}
