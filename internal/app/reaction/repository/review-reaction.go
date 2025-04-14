package repository

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/entity"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/doug-martin/goqu/v9"
)

func (r *reviewReactionRepositoryImplPostgre) CreateReaction(ctx context.Context, entity *entity.ReviewReaction) error {
	qb := goqu. 
		Insert(reviewReactionTableName). 
		Rows(entity).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()

	if err != nil {
		return apperr.NewFromError(err, "Failed to create review reaction").SetLocation()
	}

	query = r.conn.Rebind(query)

	_, err = r.conn.QueryxContext(ctx, query, args...)

	if err != nil {
		return apperr.NewFromError(err, "Failed to create review reaction").SetLocation()
	}

	return nil
}