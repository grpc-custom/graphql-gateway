package service

import (
	"context"
	"fmt"

	"github.com/grpc-custom/graphql-gateway/example/federation/proto/review"
)

type Review struct {
	ID       string
	AuthorID string
	UPC      string
	Body     string
}

var (
	reviews = []*Review{
		{
			ID:       "1",
			AuthorID: "1",
			UPC:      "1",
			Body:     "Love it!",
		},
		{
			ID:       "2",
			AuthorID: "1",
			UPC:      "2",
			Body:     "Too expensive.",
		},
		{
			ID:       "3",
			AuthorID: "2",
			UPC:      "3",
			Body:     "Could be better.",
		},
		{
			ID:       "4",
			AuthorID: "2",
			UPC:      "1",
			Body:     "Prefer something else.",
		},
	}
)

type ReviewService struct{}

func NewReviewService() review.ReviewServiceServer {
	svc := &ReviewService{}
	return svc
}

func (r *ReviewService) ListUserReviews(ctx context.Context, req *review.ListUserReviewsRequest) (*review.ListUserReviewsResponse, error) {
	res := &review.ListUserReviewsResponse{
		Reviews: []*review.Review{},
	}
	fmt.Println(req)
	for _, data := range reviews {
		if data.AuthorID == req.UserId {
			res.Reviews = append(res.Reviews, &review.Review{
				Id:        data.ID,
				AuthorId:  data.AuthorID,
				ProductId: data.UPC,
				Body:      data.Body,
			})
		}
	}
	return res, nil
}

func (r *ReviewService) ListProductReviews(ctx context.Context, req *review.ListProductReviewsRequest) (*review.ListProductReviewsResponse, error) {
	res := &review.ListProductReviewsResponse{
		Reviews: []*review.Review{},
	}
	for _, data := range reviews {
		if data.UPC == req.ProductId {
			res.Reviews = append(res.Reviews, &review.Review{
				Id:        data.ID,
				AuthorId:  data.AuthorID,
				ProductId: data.UPC,
				Body:      data.Body,
			})
		}
	}
	return res, nil
}
