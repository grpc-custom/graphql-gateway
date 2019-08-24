package service

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/grpc-custom/graphql-gateway/test/sample/proto/red"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RedService struct {
	data  sync.Map
	total int32
}

func NewReadService() red.RedServiceServer {
	return &RedService{
		data: sync.Map{},
	}
}

func (r *RedService) Get(ctx context.Context, req *red.GetRequest) (*red.GetResponse, error) {
	value, ok := r.data.Load(req.Id)
	resp := &red.GetResponse{}
	if ok {
		resp.Red = value.(*red.Red)
	}
	return resp, nil
}

func (r *RedService) List(ctx context.Context, req *red.ListRequest) (*red.ListResponse, error) {
	var list []*red.Red
	var n int32 = 1
	r.data.Range(func(key, value interface{}) bool {
		red := value.(*red.Red)
		list = append(list, red)
		n++
		return req.Limit == 0 || n <= req.Limit
	})
	resp := &red.ListResponse{
		Reds:  list,
		Limit: req.Limit,
		Size:  n,
		Total: atomic.LoadInt32(&r.total),
	}
	return resp, nil
}

func (r *RedService) Create(ctx context.Context, req *red.CreateRequest) (*red.CreateResponse, error) {
	now := time.Now()
	value := &red.Red{
		Id:        req.Id,
		Name:      req.Name,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		Enabled:   true,
		Point:     1000,
	}
	r.data.Store(req.Id, value)
	atomic.AddInt32(&r.total, 1)
	resp := &red.CreateResponse{
		Red: value,
	}
	return resp, nil
}

func (r *RedService) Update(ctx context.Context, req *red.UpdateRequest) (*red.UpdateResponse, error) {
	value, ok := r.data.Load(req.Id)
	if !ok {
		return nil, status.Error(codes.NotFound, "not found")
	}
	rr := value.(*red.Red)
	rr.Name = req.Name
	r.data.Store(req.Id, rr)
	resp := &red.UpdateResponse{
		Red: rr,
	}
	return resp, nil
}

func (r *RedService) Delete(ctx context.Context, req *red.DeleteRequest) (*red.DeleteResponse, error) {
	_, ok := r.data.Load(req.Id)
	if !ok {
		return &red.DeleteResponse{}, nil
	}
	r.data.Delete(req.Id)
	atomic.AddInt32(&r.total, -1)
	resp := &red.DeleteResponse{
		Id: req.Id,
	}
	return resp, nil
}
