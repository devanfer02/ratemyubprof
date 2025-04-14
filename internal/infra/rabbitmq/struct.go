package rabbitmq

import (
	"fmt"

	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

const (
	ReactionReviewCreateQueue = "review-reactions.create"
	ReactionReviewDeleteQueue = "review-reactions.create"
)

type RabbitMQ struct {
	logger *zap.Logger
	conn *amqp091.Connection
}

func NewRabbitMQ(env *env.Env, logger *zap.Logger) *RabbitMQ {
	conn, err := amqp091.Dial(
		fmt.Sprintf("ampq://%s:%s@%s:%d/",
			env.RabbitMQ.User,
			env.RabbitMQ.Password,
			env.RabbitMQ.Host,
			env.RabbitMQ.Port,
		),
	)
	if err != nil {
		panic(err)
	}

	rabbitMQ := &RabbitMQ{
		conn: conn,
		logger: logger,
	}

	return rabbitMQ
}

func (r *RabbitMQ) Close() {
	if err := r.conn.Close(); err != nil {
		panic(err)
	}
}