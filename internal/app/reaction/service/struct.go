package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/reaction/contracts"
	"github.com/devanfer02/ratemyubprof/internal/infra/rabbitmq"
	"go.uber.org/zap"
)

type ReviewReactionService struct {
	reactionRepo contracts.ReviewReactionRepositoryProvider
	logger *zap.Logger
	rabbitMQ *rabbitmq.RabbitMQ 
}

func NewReviewReactionService(reactionRepo contracts.ReviewReactionRepositoryProvider, logger *zap.Logger, rabbitMQ *rabbitmq.RabbitMQ) contracts.ReviewReactionService {
	return &ReviewReactionService{
		reactionRepo: reactionRepo,
		logger: logger,
		rabbitMQ:     rabbitMQ,
	}
}