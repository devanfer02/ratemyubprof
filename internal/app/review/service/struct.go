package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/review/contracts"
)

type reviewService struct {
	reviewRepo contracts.ReviewRepositoryProvider
}

func NewReviewService(reviewRepo contracts.ReviewRepositoryProvider) contracts.ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
	}
}
