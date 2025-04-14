package rabbitmq

import "github.com/rabbitmq/amqp091-go"

type Queue struct {
	Name string 
	Durable bool
	AutoDelete bool
	Exclusive bool
	NoWait bool
	Args amqp091.Table
}

func (r *RabbitMQ) DeclareQueues() error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err 
	}

	defer ch.Close()

	queues := []Queue{
		{
			Name: ReactionReviewCreateQueue,
			Durable: true,
			AutoDelete: false,
			Exclusive: true,
			NoWait: true,
			Args: nil,
		},
		{
			Name: ReactionReviewDeleteQueue,
			Durable: true,
			AutoDelete: false,
			Exclusive: true,
			NoWait: true,
			Args: nil,
		},
	}

	for _, q := range queues {
		_,  err := ch.QueueDeclare(
			q.Name,
			q.Durable,
			q.AutoDelete,
			q.Exclusive,
			q.NoWait,
			q.Args,
		); 
		if err != nil {
			return err 
		}
	}

	return nil 
}