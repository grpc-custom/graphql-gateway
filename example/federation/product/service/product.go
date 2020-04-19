package service

import (
	"context"
	"fmt"

	"github.com/grpc-custom/graphql-gateway/example/federation/proto/product"
)

var products = []*product.Product{
	{
		Upc:   "1",
		Name:  "Table",
		Price: 899,
	},
	{
		Upc:   "2",
		Name:  "Couch",
		Price: 1299,
	},
	{
		Upc:   "3",
		Name:  "Chair",
		Price: 54,
	},
}

type ProductService struct{}

func NewProductService() product.ProductServiceServer {
	svc := &ProductService{}
	return svc
}

func (p *ProductService) TopProducts(ctx context.Context, req *product.TopProductsRequest) (*product.TopProductsResponse, error) {
	fmt.Println("TopProducts --->", req)
	res := &product.TopProductsResponse{
		Products: products,
	}
	return res, nil
}

func (p *ProductService) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	fmt.Println("GetProduct --->", req)
	res := &product.GetProductResponse{}
	for i := range products {
		if products[i].Upc == req.Upc {
			res.Products = products[i]
			break
		}
	}
	return res, nil
}
