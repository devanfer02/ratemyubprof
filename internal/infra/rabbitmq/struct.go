package rabbitmq

import (
	"fmt"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)


type RabbitMQ struct {
	logger *zap.Logger
	conn *amqp091.Connection
}

func NewRabbitMQ(env *env.Env, logger *zap.Logger) *RabbitMQ {
	var (
		retries = 5
		conn *amqp091.Connection
		err error 
	)
	
	for {
		conn, err = amqp091.Dial(
			fmt.Sprintf("amqp://%s:%s@%s:%d/",
				env.RabbitMQ.User,
				env.RabbitMQ.Password,
				env.RabbitMQ.Host,
				env.RabbitMQ.Port,
			),
		)
		if err == nil {
			break 
		}

		if (retries <= 0) {
			logger.Error("Failed to connect to RabbitMQ after multiple attempts", zap.Error(err))
			panic(err)
		}

		logger.Error("Failed to connect to RabbitMQ", zap.Error(err))
		logger.Info("Retrying in 5 seconds...")

		time.Sleep(5 * time.Second)
		retries--
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