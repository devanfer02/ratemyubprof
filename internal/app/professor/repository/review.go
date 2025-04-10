package repository

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
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
		return apperr.NewFromError(err, "Failed to insert review professor").SetLocation()
	}

	query = p.conn.Rebind(query)

	_, err = p.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return apperr.NewFromError(err, "Failed to insert review professor").SetLocation()
	}

	return nil
}

func (p *professorRepositoryImplPostgre) DeleteProfessorReview(ctx context.Context, params *dto.FetchReviewParams) error {
	qb := goqu. 
		Delete(reviewTableName). 
		Where(
			goqu.And(goqu.I("reviews.prof_id").Eq(params.ProfId), 
			goqu.I("reviews.user_id").Eq(params.UserId)),
		)

	query, args, err := qb.ToSQL()
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review professor").SetLocation()
	}

	query = p.conn.Rebind(query)

	_, err = p.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review professor").SetLocation()
	}

	return nil 
}