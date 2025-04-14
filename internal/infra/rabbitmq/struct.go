package rabbitmq

import (
	"fmt"

	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)


type RabbitMQ struct {
	logger *zap.Logger
	conn *amqp091.Connection
}

func NewRabbitMQ(env *env.Env, logger *zap.Logger) *RabbitMQ {
	conn, err := amqp091.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%d/",
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

func (r *RabbitMQ) Close() error {
	return r.conn.Close()
}