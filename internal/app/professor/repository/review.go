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


func (p *professorRepositoryImplPostgre) GetReviewsItemsByProfID(ctx context.Context, profId string) (uint64, error) {
	var count uint64
	
	qb := goqu.
		From(reviewTableName).
		Select(goqu.COUNT(goqu.Star())).
		Where(goqu.I("reviews.id").Eq(profId)).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return 0, err
	}

	query = p.conn.Rebind(query)

	err = p.conn.QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}