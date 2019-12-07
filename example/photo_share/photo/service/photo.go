package service

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-custom/graphql-gateway/example/photo_share/proto/photo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type PhotoService struct {
	data sync.Map
}

func NewPhotoService() photo.PhotoServiceServer {
	svc := &PhotoService{
		data: sync.Map{},
	}
	svc.data.Store("1", &photo.Photo{
		Id:          "1",
		Name:        "sample-1",
		Category:    photo.PhotoCategory_GRAPHIC,
		Description: "sample-description-1",
		Url:         "http://localhost/photos/1",
		Created:     ptypes.TimestampNow(),
	})
	svc.data.Store("2", &photo.Photo{
		Id:          "2",
		Name:        "sample-2",
		Category:    photo.PhotoCategory_PORTRAIT,
		Description: "sample-description-2",
		Url:         "http://localhost/photos/2",
		Created:     ptypes.TimestampNow(),
	})
	svc.data.Store("3", &photo.Photo{
		Id:          "3",
		Name:        "sample-3",
		Category:    photo.PhotoCategory_ACTION,
		Description: "sample-description-3",
		Url:         "http://localhost/photos/3",
		Created:     ptypes.TimestampNow(),
	})
	return svc
}

func (p *PhotoService) TotalPhotos(ctx context.Context, _ *empty.Empty) (*photo.TotalPhotosResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md.Get("Authorization"))
		fmt.Println(md.Get("X-Forwarded-For"))
		fmt.Println(md.Get("graphqlgateway-user-agent"))
	}
	var total int32
	p.data.Range(func(_, _ interface{}) bool {
		total++
		return true
	})
	resp := &photo.TotalPhotosResponse{
		Total: total,
	}
	return resp, nil
}

func (p *PhotoService) AllPhotos(ctx context.Context, _ *empty.Empty) (*photo.AllPhotosResponse, error) {
	var photos []*photo.Photo
	p.data.Range(func(_, value interface{}) bool {
		data, ok := value.(*photo.Photo)
		if !ok {
			return true
		}
		photos = append(photos, data)
		return true
	})
	resp := &photo.AllPhotosResponse{
		Photos: photos,
	}
	return resp, nil
}

func (p *PhotoService) Photo(ctx context.Context, in *photo.PhotoRequest) (*photo.PhotoResponse, error) {
	if in.Id == "error" {
		err := status.New(codes.InvalidArgument, "some error")
		detail := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "id",
					Description: "invalid format",
				},
			},
		}
		dt, _ := err.WithDetails(detail)
		return nil, dt.Err()
	}
	value, ok := p.data.Load(in.Id)
	if !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no such photo data: %s", in.Id))
	}
	resp := &photo.PhotoResponse{
		Photo: value.(*photo.Photo),
	}
	return resp, nil
}

func (p *PhotoService) PostPhoto(ctx context.Context, in *photo.PostPhotoRequest) (*photo.PhotoResponse, error) {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	data := &photo.Photo{
		Id:          id,
		Name:        in.Name,
		Description: in.Description,
		Category:    in.Category,
		Url:         fmt.Sprintf("http://localhost/photos/%s", id),
		Created:     ptypes.TimestampNow(),
	}
	p.data.Store(id, data)
	resp := &photo.PhotoResponse{
		Photo: data,
	}
	return resp, nil
}

func (p *PhotoService) TagPhoto(ctx context.Context, in *photo.TagPhotoRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
