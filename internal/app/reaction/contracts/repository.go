package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type ReviewReactionRepositoryProvider interface {
	NewClient(tx bool) (ReviewReactionRepository, error)
}

type ReviewReactionRepository interface {
	FetchReactionByParams(ctx context.Context, params *dto.FetchReactionParams) (entity.ReviewReaction, error)
	CreateReaction(ctx context.Context, entity *entity.ReviewReaction) error
	DeleteReaction(ctx context.Context, entity *entity.ReviewReaction) error
}