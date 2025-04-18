package service

import (
	"context"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/reaction/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/devanfer02/ratemyubprof/internal/infra/rabbitmq"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
)

func (s *reviewReactionService) PublishReaction(ctx context.Context, queueType rabbitmq.QueueType, req *dto.ReviewReactionRequest) error {
	// Publish to RabbitMQ
	err := s.rabbitMQ.Publish(ctx, queueType, req)

	if err != nil {
		return apperr.NewFromError(err, "Failed to publish review reaction").SetLocation()
	}

	return nil 
}

func (s *reviewReactionService) CreateReaction(ctx context.Context, req *dto.ReviewReactionRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	repoClient, err := s.reactionRepo.NewClient(false)
	if err != nil {
		return err 
	}

	err = repoClient.DeleteReaction(ctx, &entity.ReviewReaction{UserID: req.UserID, ReviewID: req.ReviewID})

	if err != nil && err != contracts.ErrItemNotFound {
		return err 
	}

	entity := formatter.FormatReactionToEntity(req)
	err = repoClient.CreateReaction(ctx, entity)

	if err != nil {
		return apperr.NewFromError(err, "Failed to create review reaction").SetLocation()
	}

	select {
	case <- ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return nil 
	}

}

func (s *reviewReactionService) DeleteReaction(ctx context.Context, req *dto.ReviewReactionRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	repoClient, err := s.reactionRepo.NewClient(false)
	if err != nil {
		return err 
	}

	entity := formatter.FormatReactionToEntity(req)
	err = repoClient.DeleteReaction(ctx, entity)

	if err != nil {
		return apperr.NewFromError(err, "Failed to create review reaction").SetLocation()
	}

	select {
	case <- ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return nil 
	}
}