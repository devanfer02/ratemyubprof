package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type ReviewRepositoryProvider interface {
	NewClient(tx bool) (ReviewRepository, error)
}

type ReviewRepository interface {
	FetchReviewsByParams(ctx context.Context, params *dto.FetchReviewParams, pageQuery *dto.PaginationQuery) ([]entity.ReviewWithRelations, error)
	GetReviewsItemsByParams(ctx context.Context, params *dto.FetchReviewParams) (uint64, error)
	FetchRatingDistributionByProfId(ctx context.Context, profId string, column entity.RatingDistributionCol) (entity.RatingDistribution, error)
}