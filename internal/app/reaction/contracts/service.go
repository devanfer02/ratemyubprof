package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/infra/rabbitmq"
)

type ReviewReactionService interface {
	PublishReaction(ctx context.Context, queueType rabbitmq.QueueType, req *dto.ReviewReactionRequest) error
	CreateReaction(ctx context.Context, req *dto.ReviewReactionRequest) error
}