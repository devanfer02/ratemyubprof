package repository

import (
	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/jmoiron/sqlx"
)

var (
	userTableName = "users"
)

type repository struct {
	conn *sqlx.DB
}

type userRepositoryImplPostgre struct {
	conn sqlx.ExtContext
}

func NewUserRepository(conn *sqlx.DB) contracts.UserRepositoryProvider {
	return &repository{
		conn: conn,
	}
}

func (r *repository) NewClient(tx bool) (contracts.UserRepository, error) {
	var (
		conn  sqlx.ExtContext
		err error
	)

	if tx {
		conn, err = r.conn.Beginx()
		if err != nil {
			return nil, err
		}
	} else {
		conn = r.conn
	}

	return &userRepositoryImplPostgre{
		conn: conn,
	}, nil
}