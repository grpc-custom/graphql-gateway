package service

import (
	"context"
	"fmt"

	"github.com/grpc-custom/graphql-gateway/example/mixin/proto/review"
)

type reviewService struct{}

func New() review.ReviewServiceServer {
	return &reviewService{}
}

func (r *reviewService) GetReview(ctx context.Context, req *review.GetReviewRequest) (*review.GetReviewResponse, error) {
	res := &review.GetReviewResponse{
		Review: &review.Review{
			Id:         "test",
			Body:       "test",
			AuthorId:   "test",
			ProductUpc: "test",
		},
	}
	return res, nil
}

func (r *reviewService) ListUserReviews(ctx context.Context, req *review.ListUserReviewsRequest) (*review.ListUserReviewsResponse, error) {
	fmt.Println("ListUserReviews", req)
	res := &review.ListUserReviewsResponse{
		Reviews: []*review.Review{
			{
				Id:         req.Id,
				Body:       "body-" + req.Id,
				AuthorId:   "author-" + req.Id,
				ProductUpc: "product-" + req.Id,
			},
		},
	}
	return res, nil
}

func (r *reviewService) ListProductReviews(ctx context.Context, req *review.ListProductReviewsRequest) (*review.ListProductReviewsResponse, error) {
	fmt.Println("ListProductReviews", req)
	res := &review.ListProductReviewsResponse{
		Reviews: []*review.Review{
			{
				Id:         req.Upc,
				Body:       "body-" + req.Upc,
				AuthorId:   "author-" + req.Upc,
				ProductUpc: "product-" + req.Upc,
			},
		},
	}
	return res, nil
}
