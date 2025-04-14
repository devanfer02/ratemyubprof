package rabbitmq

import (
	"context"

	"github.com/bytedance/sonic"
	"go.uber.org/zap"
)

func Consume[T any](ctx context.Context, queueName string, rabbit *RabbitMQ) (<- chan T, error) {
	ch, err := rabbit.conn.Channel()
	if err != nil {
		return nil, err 
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		queueName,
		"ratemyubprof.internal.consumer",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err 
	}

	out := make(chan T)

	go func() {
		defer ch.Close()
		defer close(out)
		for d := range msgs {
			var data T 

			if err := sonic.Unmarshal(d.Body, &data); err != nil {
				rabbit.logger.Error(
					"[RabbitMQ] Error Unmarshalling Message",
					zap.String("Queue", queueName),
					zap.String("Message", string(d.Body)),
					zap.String("Error", err.Error()),
				)
				continue 
			}

			out <- data
		}
	}()
	

	return out, nil 
}