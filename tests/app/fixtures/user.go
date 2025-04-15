package fixtures

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/golang-migrate/migrate/v4"
	mp "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

type CleanUpFunc = func()

func NewDB() (*sqlx.DB, CleanUpFunc) {
	ctx := context.Background()

	pgC, err := postgres.Run(
		ctx,
		"postgres:16.1",
		postgres.WithDatabase("ratemyubproftest"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2). 
				WithStartupTimeout(5 * time.Second),
		),
	)

	if err != nil {
		panic(err)
	}

	connStr, err := pgC.ConnectionString(ctx, "sslmode=disable")

	if err != nil {
		panic(err)
	}
	
	dbx := sqlx.MustConnect("postgres", connStr)

	driver := must(mp.WithInstance(dbx.DB, &mp.Config{}))


	m := must(migrate.NewWithDatabaseInstance(
		"file://../../../internal/infra/database/postgres/migrations",
		"ratemyubproftest", driver,
	))

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	return dbx, func() {
		pgC.Terminate(ctx)
		dbx.Close()
	}
}
