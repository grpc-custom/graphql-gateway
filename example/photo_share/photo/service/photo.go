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
	"google.golang.org/grpc/codes"
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
	var total int
	p.data.Range(func(_, _ interface{}) bool {
		total++
		return true
	})
	resp := &photo.TotalPhotosResponse{
		Total: int32(total),
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

func (p *PhotoService) TagPhoto(ctx context.Context, in *photo.TagPhotoRequest) (*photo.PhotoResponse, error) {
	return nil, nil
}
