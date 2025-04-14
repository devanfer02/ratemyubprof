package service

import (
	"context"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/reaction/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/infra/rabbitmq"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
	"go.uber.org/zap"
)

type Worker struct {
	QueueType rabbitmq.QueueType
	HandleFn func(ctx context.Context, req *dto.ReviewReactionRequest) error
}

func (s *ReviewReactionService) PublishReaction(ctx context.Context, queueType rabbitmq.QueueType, req *dto.ReviewReactionRequest) error {
	// Publish to RabbitMQ
	err := s.rabbitMQ.Publish(ctx, queueType, req)

	if err != nil {
		return apperr.NewFromError(err, "Failed to publish review reaction").SetLocation()
	}

	return nil 
}

func (s *ReviewReactionService) CreateReaction(ctx context.Context, req *dto.ReviewReactionRequest) error {
	// Spawn context with timeout
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	// Insert Creation to Database
	repoClient, err := s.reactionRepo.NewClient(false)
	if err != nil {
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

func (s *ReviewReactionService) StartWorkers(ctx context.Context) {
	actions := []Worker{
		{
			QueueType: rabbitmq.ReactionReviewCreateQueue,
			HandleFn: s.CreateReaction,
		},
	}

	for _, action := range actions {
		go func(worker Worker) {			
			reqs, err := rabbitmq.Consume[dto.ReviewReactionRequest](
				ctx,
				worker.QueueType.String(),
				s.rabbitMQ,
			)

			if err != nil {
				s.logger.Error(
					"[RabbitMQ] Error consuming message",
					zap.String("Queue", worker.QueueType.String()),
					zap.String("Error", err.Error()),
				)
				return 
			}

			// Limit to 10 concurrent actions per worker 
			sem := make(chan struct{}, 10)

			for req := range reqs {
				s.logger.Info(
					"[RabbitMQ] Received message!",
					zap.String("Queue", worker.QueueType.String()),
				)

				sem <- struct{}{}

				go func(req dto.ReviewReactionRequest) {
					defer func() { <- sem }()

					if err := action.HandleFn(ctx, &req); err != nil {
						s.logger.Error(
							"[RabbitMQ] Error handling message",
							zap.String("Queue", worker.QueueType.String()),
							zap.String("Error", err.Error()),
						)
					}
				}(req)

			}
		}(action)
	}
}

