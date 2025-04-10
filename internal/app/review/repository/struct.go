package repository

import (
	"github.com/devanfer02/ratemyubprof/internal/app/review/contracts"
	"github.com/jmoiron/sqlx"
)

var (
	reviewTableName    = "reviews"
	professorTableName = "professors"
	userTableName      = "users"
)

type repository struct {
	conn *sqlx.DB
}

type reviewRepositoryImplPostgre struct {
	conn sqlx.ExtContext
}

func NewReviewRepository(conn *sqlx.DB) contracts.ReviewRepositoryProvider {
	return &repository{
		conn: conn,
	}
}

func (r *repository) NewClient(tx bool) (contracts.ReviewRepository, error) {
	var (
		conn sqlx.ExtContext
		err  error
	)

	if tx {
		conn, err = r.conn.Beginx()
		if err != nil {
			return nil, err
		}
	} else {
		conn = r.conn
	}

	return &reviewRepositoryImplPostgre{
		conn: conn,
	}, nil
}