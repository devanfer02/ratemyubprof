package repository

import (
	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/jmoiron/sqlx"
)

var (
	professorTableName = "professors"
	reviewTableName    = "reviews"
)

type repository struct {
	conn *sqlx.DB
}

type professorRepositoryImplPostgre struct {
	conn sqlx.ExtContext
}

func NewProfessorRepository(conn *sqlx.DB) contracts.ProfessorRepositoryProvider {
	return &repository{
		conn: conn,
	}
}

func (r *repository) NewClient(tx bool) (contracts.ProfessorRepository, error) {
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

	return &professorRepositoryImplPostgre{
		conn: conn,
	}, nil
}