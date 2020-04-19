package service

import (
	"context"
	"fmt"

	"github.com/grpc-custom/graphql-gateway/example/federation/proto/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type User struct {
	ID        string
	Name      string
	BirthDate string
	Username  string
}

var (
	users = []*User{
		{
			ID:        "1",
			Name:      "Ada Lovelace",
			BirthDate: "1815-12-10",
			Username:  "@ada",
		},
		{
			ID:        "2",
			Name:      "Alan Turing",
			BirthDate: "1912-06-23",
			Username:  "@complete",
		},
	}
)

type AccountService struct{}

func NewAccountService() account.AccountServiceServer {
	svc := &AccountService{}
	return svc
}

func (a *AccountService) GetMe(ctx context.Context, req *account.GetMeRequest) (*account.GetMeResponse, error) {
	var id string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md.Get("Authorization")) > 0 {
		id = md.Get("Authorization")[0]
	}
	fmt.Println(id)
	res := &account.GetMeResponse{}
	for i := range users {
		if users[i].ID == id {
			res.User = &account.User{
				Id:       users[i].ID,
				Name:     users[i].Name,
				Username: users[i].Username,
			}
		}
	}
	return res, nil
}

func (a *AccountService) GetUser(ctx context.Context, req *account.GetUserRequest) (*account.GetUserResponse, error) {
	fmt.Println("GetUser -->", req)
	for i := range users {
		if users[i].ID == req.Id {
			res := &account.GetUserResponse{
				User: &account.User{
					Id:       users[i].ID,
					Name:     users[i].Name,
					Username: users[i].Username,
				},
			}
			return res, nil
		}
	}
	return nil, status.Error(codes.NotFound, "account: user not found")
}

func (a *AccountService) MultiGetUsers(ctx context.Context, req *account.MultiGetUsersRequest) (*account.MultiGetUsersResponse, error) {
	fmt.Println("MultiGetUsers -->", req)
	ids := make(map[string]struct{}, len(req.Ids))
	for i := range req.Ids {
		ids[req.Ids[i]] = struct{}{}
	}
	res := &account.MultiGetUsersResponse{
		Users: make([]*account.User, 0, len(req.Ids)),
	}
	for i := range users {
		if _, ok := ids[users[i].ID]; ok {
			res.Users = append(res.Users, &account.User{
				Id:       users[i].ID,
				Name:     users[i].Name,
				Username: users[i].Username,
			})
		}
	}
	return res, nil
}
