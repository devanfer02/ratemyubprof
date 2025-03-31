package repository

import (
	"context"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/doug-martin/goqu/v9"
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

func (p *professorRepositoryImplPostgre) InsertProfessorsBulk(ctx context.Context, professors []entity.Professor) error {
	records := make([]goqu.Record, len(professors))
	for i, d := range professors {
		records[i] = goqu.Record{
			"id":               d.ID,
			"name":             d.Name,
			"faculty":          d.Faculty,
			"major":            d.Major,
			"profile_img_link": d.ProfileImgLink,
			"created_at":       time.Now(),
			"updated_at":       time.Now(),
		}
	}

	qb := goqu.
		Insert(professorTableName).
		Rows(records).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return err
	}

	query = p.conn.Rebind(query)

	_, err = p.conn.ExecContext(ctx, query, args...)
	if err != nil {

		return err
	}

	return nil
}

func (p *professorRepositoryImplPostgre) InsertProfessorReview(ctx context.Context, review *entity.Review) error {
	qb := goqu.
		Insert(reviewTableName).
		Rows(review).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return err
	}

	query = p.conn.Rebind(query)

	_, err = p.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
