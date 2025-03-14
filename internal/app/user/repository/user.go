package repository

import (
	contracts "github.com/devanfer02/presentia-api/internal/contracts/repositories"
	"github.com/jmoiron/sqlx"
)

type userRepositoryImplPostgre struct {
	conn sqlx.ExtContext
}

func NewUserRepository(conn *sqlx.DB) contracts.UserRepository {
	return &userRepositoryImplPostgre{
		conn: conn,
	}
}