package service

import (
	"context"

	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/product"
)

type productService struct{}

func New() product.ProductServiceServer {
	return &productService{}
}

func (p *productService) TopProducts(ctx context.Context, req *product.TopProductsRequest) (*product.TopProductsResponse, error) {
	res := &product.TopProductsResponse{
		Products: []*product.Product{
			{
				Upc:    "test-01",
				Name:   "test-01",
				Price:  100,
				Weight: 10,
			},
			{
				Upc:    "test-02",
				Name:   "test-02",
				Price:  200,
				Weight: 20,
			},
		},
	}
	return res, nil
}

func (p *productService) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	res := &product.GetProductResponse{
		Product: &product.Product{
			Upc:    "test",
			Name:   "test",
			Price:  100,
			Weight: 10,
		},
	}
	return res, nil
}
