package database

import (
	"fmt"

	"github.com/devanfer02/presentia-api/internal/infra/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(env *env.Env) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.Database.Host,
		env.Database.Port,
		env.Database.User,
		env.Database.Password,
		env.Database.Name,
		env.Database.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	client := db.DB
	client.SetMaxIdleConns(10)
	client.SetMaxOpenConns(10)

	return db
}
