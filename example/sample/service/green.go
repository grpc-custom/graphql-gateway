package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/grpc-custom/graphql-gateway/example/sample/proto/green"
)

type GreenService struct{}

func NewGreenService() green.GreenServiceServer {
	return &GreenService{}
}

func fakeGreen() *green.Green {
	now := time.Now()
	sec := now.Unix()
	nano := int32(now.Sub(time.Unix(sec, 0)))
	return &green.Green{
		Id:        "green",
		Title:     "title",
		Text:      "text, text, text",
		Url:       "http://example.com/sample",
		CreatedAt: &timestamp.Timestamp{Seconds: sec, Nanos: nano},
		UpdatedAt: &timestamp.Timestamp{Seconds: sec, Nanos: nano},
		Term: &green.Term{
			StartAt: now.Unix(),
			EndAt:   now.AddDate(0, 0, 1).Unix(),
		},
		Published: true,
	}
}

func (g *GreenService) List(ctx context.Context, req *green.ListRequest) (*green.ListResponse, error) {
	list := make([]*green.Green, 10)
	for i := range list {
		list[i] = fakeGreen()
	}
	resp := &green.ListResponse{
		Greens: list,
	}
	return resp, nil
}

func (g *GreenService) Get(ctx context.Context, req *green.GetRequest) (*green.GetResponse, error) {
	resp := &green.GetResponse{
		Green: fakeGreen(),
	}
	return resp, nil
}
