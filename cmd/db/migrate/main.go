package main

import (
	"github.com/devanfer02/ratemyubprof/internal/infra/database"
	"github.com/devanfer02/ratemyubprof/internal/infra/env"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	env := env.NewEnv()
	db := database.NewDatabase(env)

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/infra/database/migrations",
		env.Database.Name, driver,
	)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

}