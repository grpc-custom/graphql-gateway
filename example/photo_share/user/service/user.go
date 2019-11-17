package service

import (
	"context"
	"strings"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-custom/graphql-gateway/example/photo_share/proto/user"
	"google.golang.org/grpc/metadata"
)

type UserService struct {
	data sync.Map
}

func NewUserService() user.UserServerServer {
	svc := &UserService{
		data: sync.Map{},
	}
	svc.data.Store("test", &user.User{
		Name:   "test-name",
		Avatar: "test-avatar",
	})
	return svc
}

func (u *UserService) Me(ctx context.Context, _ *empty.Empty) (*user.UserResponse, error) {
	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md.Get("Authorization")) > 0 {
		auth := md.Get("Authorization")[0]
		arr := strings.SplitN(auth, " ", 2)
		token = arr[1]
	}
	resp := &user.UserResponse{}
	if data, ok := u.data.Load(token); ok {
		resp.User = data.(*user.User)
	}
	return resp, nil
}

func (u *UserService) TotalUsers(ctx context.Context, _ *empty.Empty) (*user.TotalUsersResponse, error) {
	var total int32
	u.data.Range(func(_, _ interface{}) bool {
		total++
		return true
	})
	resp := &user.TotalUsersResponse{
		Total: total,
	}
	return resp, nil
}

func (u *UserService) AllUsers(ctx context.Context, _ *empty.Empty) (*user.AllUsersResponse, error) {
	var users []*user.User
	u.data.Range(func(_, value interface{}) bool {
		data, ok := value.(*user.User)
		if !ok {
			return true
		}
		users = append(users, data)
		return true
	})
	resp := &user.AllUsersResponse{
		Users: users,
	}
	return resp, nil
}

func (u *UserService) User(ctx context.Context, in *user.LoginRequest) (*user.UserResponse, error) {
	return &user.UserResponse{}, nil
}

func (u *UserService) GithubAuth(ctx context.Context, in *user.GithubAuthRequest) (*user.GithubAuthResponse, error) {
	return &user.GithubAuthResponse{}, nil
}
