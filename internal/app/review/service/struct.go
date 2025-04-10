package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/review/contracts"
	"go.uber.org/zap"
)

type reviewService struct {
	reviewRepo contracts.ReviewRepositoryProvider
	logger *zap.Logger
}

func NewReviewService(logger *zap.Logger, reviewRepo contracts.ReviewRepositoryProvider) contracts.ReviewService {
	return &reviewService{
		logger: logger,
		reviewRepo: reviewRepo,
	}
}
