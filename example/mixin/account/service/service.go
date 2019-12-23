package service

import (
	"context"
	"fmt"

	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/account"
)

type accountService struct{}

func New() account.AccountServiceServer {
	return &accountService{}
}

func (a *accountService) GetMe(ctx context.Context, req *account.GetMeRequest) (*account.GetMeResponse, error) {
	fmt.Println("GetMe")
	ret := &account.GetMeResponse{
		User: &account.User{
			Id:       "test",
			Name:     "me",
			Username: "test-me",
		},
	}
	return ret, nil
}

func (a *accountService) GetUser(ctx context.Context, req *account.GetUserRequest) (*account.GetUserResponse, error) {
	ret := &account.GetUserResponse{
		User: &account.User{
			Id:       req.Id,
			Name:     "name-" + req.Id,
			Username: "username-" + req.Id,
		},
	}
	return ret, nil
}
