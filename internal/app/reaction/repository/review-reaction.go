package repository

import (
	"context"
	"database/sql"

	"github.com/devanfer02/ratemyubprof/internal/app/reaction/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/doug-martin/goqu/v9"
)

func (r *reviewReactionRepositoryImplPostgre) FetchReactionByParams(ctx context.Context, params *dto.FetchReactionParams) (entity.ReviewReaction, error) {
	qb := goqu. 
		From(goqu.T(reviewReactionTableName).As("rr")).
		Select("*"). 
		Where(
			goqu.And(
				goqu.I("rr.user_id").Eq(params.UserID),
				goqu.I("rr.review_id").Eq(params.ReviewID),
			), 
		). 
		SetDialect(goqu.GetDialect("postgres")). 
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return entity.ReviewReaction{}, apperr.NewFromError(err, "Failed to fetch review reaction").SetLocation()
	}

	query = r.conn.Rebind(query)

	var res entity.ReviewReaction

	err = r.conn.QueryRowxContext(ctx, query, args...).StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.ReviewReaction{}, contracts.ErrItemNotFound
		}
		return entity.ReviewReaction{}, apperr.NewFromError(err, "Failed to fetch review reaction").SetLocation()
	}

	return res, nil 
}

func (r *reviewReactionRepositoryImplPostgre) CreateReaction(ctx context.Context, entity *entity.ReviewReaction) error {
	qb := goqu.
		Insert(reviewReactionTableName).
		Rows(goqu.Record{
			"review_id":     entity.ReviewID,
			"user_id":       entity.UserID,
			"reaction_type": entity.Type,
		}).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()

	if err != nil {
		return apperr.NewFromError(err, "Failed to create review reaction").SetLocation()
	}

	query = r.conn.Rebind(query)

	_, err = r.conn.QueryxContext(ctx, query, args...)

	if err != nil {

		if contracts.IsErrorCode(err, contracts.PgsqlUniqueViolationErr) {
			return contracts.ErrItemAlreadyExists
		}

		return apperr.NewFromError(err, "Failed to create review reaction").SetLocation()
	}

	return nil
}

func (r *reviewReactionRepositoryImplPostgre) DeleteReaction(ctx context.Context, entity *entity.ReviewReaction) error {
	qb := goqu. 
		Delete(reviewReactionTableName). 
		Where(goqu.And(
			goqu.Ex{"review_id": entity.ReviewID},
			goqu.Ex{"user_id": entity.UserID},
		)). 
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review reaction").SetLocation()
	}

	query = r.conn.Rebind(query)

	res, err := r.conn.ExecContext(ctx, query, args...)

	if err != nil {
		return apperr.NewFromError(err, "Failed to delete review reaction").SetLocation()
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