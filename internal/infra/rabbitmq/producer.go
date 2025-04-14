package rabbitmq

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQ) Publish(ctx context.Context, queueName QueueType, data any) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err 
	}

	defer ch.Close()

	body, err := sonic.Marshal(data)
	if err != nil {
		return err 
	}

	return ch.PublishWithContext(
		ctx,
		"",
		queueName.String(),
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,		
		},
	)
}