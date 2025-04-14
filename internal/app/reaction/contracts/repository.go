package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type ReviewReactionRepositoryProvider interface {
	NewClient(tx bool) (ReviewReactionRepository, error)
}

type ReviewReactionRepository interface {
	CreateReaction(ctx context.Context, entity *entity.ReviewReaction) error
}