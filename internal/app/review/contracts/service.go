package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
)

type ReviewService interface {
	FetchReviewsByParams(ctx context.Context, params *dto.FetchReviewParams, pageQuery *dto.PaginationQuery) ([]dto.FetchReviewResponse, dto.PaginationResponse, error)	
}