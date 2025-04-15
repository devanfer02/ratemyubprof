package rabbitmq

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"go.uber.org/zap"
)

type ReactionWorker struct {
	QueueType QueueType
	HandleFn func(ctx context.Context, req *dto.ReviewReactionRequest) error
}


func (r *RabbitMQ) StartReactionWorkers(ctx context.Context ,workers []ReactionWorker) {
	for _, action := range workers {
		go func(worker ReactionWorker) {			
			reqs, err := Consume[dto.ReviewReactionRequest](
				ctx,
				worker.QueueType.String(),
				r,
			)

			if err != nil {
				r.logger.Error(
					"[RabbitMQ] Error consuming message",
					zap.String("Queue", worker.QueueType.String()),
					zap.String("Error", err.Error()),
				)
				return 
			}

			// Limit to 10 concurrent actions per worker 
			sem := make(chan struct{}, 10)

			for req := range reqs {
				r.logger.Info(
					"[RabbitMQ] Received message!",
					zap.String("Queue", worker.QueueType.String()),
				)

				sem <- struct{}{}

				go func(req dto.ReviewReactionRequest) {
					defer func() { <- sem }()

					if err := action.HandleFn(ctx, &req); err != nil {
						r.logger.Error(
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