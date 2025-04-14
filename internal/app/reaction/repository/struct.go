package repository

import (
	"github.com/devanfer02/ratemyubprof/internal/app/reaction/contracts"
	"github.com/jmoiron/sqlx"
)

var (
	reviewReactionTableName = "review_reactions"
	reviewTableName         = "reviews"
)

type repository struct {
	conn *sqlx.DB
}

type reviewReactionRepositoryImplPostgre struct {
	conn sqlx.ExtContext
}

func NewReviewReactionRepository(conn *sqlx.DB) contracts.ReviewReactionRepositoryProvider {
	return &repository{
		conn: conn,
	}
}

func (r *repository) NewClient(tx bool) (contracts.ReviewReactionRepository, error) {
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

	return &reviewReactionRepositoryImplPostgre{
		conn: conn,
	}, nil
}
