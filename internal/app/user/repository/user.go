package repository

import (
	"github.com/devanfer02/presentia-api/internal/app/user/contracts"
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