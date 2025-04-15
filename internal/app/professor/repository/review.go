package repository

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
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
		if contracts.IsErrorCode(err, contracts.PgsqlUniqueViolationErr) {
			return contracts.ErrItemAlreadyExists
		}
		
		return apperr.NewFromError(err, "Failed to insert review professor").SetLocation()
	}

	return nil
}

func (p *professorRepositoryImplPostgre) UpdateProfessorReview(ctx context.Context, review *entity.Review) error {
	qb := goqu. 
		Update(reviewTableName). 
		Set(review). 
		Where(
			goqu.And(
				goqu.I("reviews.prof_id").Eq(review.ProfessorID), 
				goqu.I("reviews.user_id").Eq(review.UserID),
			),
		)

	query, args, err := qb.ToSQL()
	if err != nil {
		return apperr.NewFromError(err, "Failed to update professor review").SetLocation()
	}

	query = p.conn.Rebind(query)

	res, err := p.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review professor").SetLocation()	
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review professor").SetLocation()	
	}

	if rows == 0 {
		return contracts.ErrItemNotFound
	}

	if rows > 1 {
		return contracts.ErrMoreThanOneAffected
	}

	return nil 
		
}

func (p *professorRepositoryImplPostgre) DeleteProfessorReview(ctx context.Context, params *dto.FetchReviewParams) error {
	qb := goqu. 
		Delete(reviewTableName). 
		Where(
			goqu.And(
				goqu.I("reviews.prof_id").Eq(params.ProfId), 
				goqu.I("reviews.user_id").Eq(params.UserId),
			),
		). 
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review professor").SetLocation()
	}

	query = p.conn.Rebind(query)

	res, err := p.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review professor").SetLocation()
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review professor").SetLocation()
	}

	if rows == 0 {
		return contracts.ErrItemNotFound
	}

	if rows > 1 {
		return contracts.ErrMoreThanOneAffected
	}

	return nil 
}