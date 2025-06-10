package fixtures

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func NewRabbitMq() (string, CleanUpFunc) {
	ctx := context.Background()

	rmC, err := rabbitmq.Run(
		ctx,
		"rabbitmq:4.0-management",
		rabbitmq.WithAdminUsername("guest"),
		rabbitmq.WithAdminPassword("guest"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("Server startup complete"),
		),
	)

	if err != nil {
		panic(err)
	}

	connStr, err := rmC.AmqpURL(ctx)

	if err != nil {
		panic(err)
	}
	
	return connStr, func() {
		rmC.Terminate(ctx)
	}
}
