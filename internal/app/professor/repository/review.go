package repository

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/doug-martin/goqu/v9"
)

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
