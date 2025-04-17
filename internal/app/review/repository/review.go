package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/devanfer02/ratemyubprof/internal/app/review/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/doug-martin/goqu/v9"
)

func (p *reviewRepositoryImplPostgre) FetchReviewsByParams(ctx context.Context, params *dto.FetchReviewParams, pageQuery *dto.PaginationQuery) ([]entity.ReviewWithRelations, error) {
	qb := goqu.
		Select(
			goqu.I("r.id").As("id"),
			goqu.I("r.prof_id").As("prof_id"),
			goqu.I("r.user_id").As("user_id"),
			goqu.I("r.comment").As("comment"),
			goqu.I("r.difficulty_rating").As("difficulty_rating"),
			goqu.I("r.friendly_rating").As("friendly_rating"),
			goqu.I("r.created_at").As("created_at"),
			goqu.I("u.id").As(goqu.C("user.id")),
			goqu.I("u.username").As(goqu.C("user.username")),
			goqu.I("p.id").As(goqu.C("professor.id")),
			goqu.I("p.name").As(goqu.C("professor.name")),
			goqu.I("p.faculty").As(goqu.C("professor.faculty")),
			goqu.I("p.major").As(goqu.C("professor.major")),
			goqu.I("p.profile_img_link").As(goqu.C("professor.profile_img_link")),
			goqu.COUNT(
				goqu.Case().
					When(goqu.I("rr.reaction_type").Eq(entity.LikeReactionType), 1). 
					Else(nil),
			).As("like_counter"),
			goqu.COUNT(
				goqu.Case().
					When(goqu.I("rr.reaction_type").Eq(entity.DislikeReactionType), 1). 
					Else(nil),
			).As("dislike_counter"),
		).
		GroupBy(goqu.I("r.id"), goqu.I("u.id"), goqu.I("p.id")).
		From(goqu.T(reviewTableName).As("r")).
		Join(
			goqu.T(userTableName).As("u"),
			goqu.On(goqu.I("r.user_id").Eq(goqu.I("u.id"))),
		).
		Join(
			goqu.T(professorTableName).As("p"),
			goqu.On(goqu.I("r.prof_id").Eq(goqu.I("p.id"))),
		). 
		LeftJoin(
			goqu.T(reactionTableName).As("rr"),
			goqu.On(goqu.I("r.id").Eq(goqu.I("rr.review_id"))),
		)	

	if params.ProfId != "" {
		qb = qb.Where(goqu.I("p.id").Eq(params.ProfId))
	}

	if params.UserId != "" {
		qb = qb.Where(goqu.I("u.id").Eq(params.UserId))
	}

	if pageQuery.Page != 0 && pageQuery.Limit != 0{
		qb = qb.Offset((pageQuery.Page - 1) * pageQuery.Limit).Limit(pageQuery.Limit)
	}

	query, args, err := qb.SetDialect(goqu.GetDialect("postgres")).Prepared(true).ToSQL()
	if err != nil {
		return nil, apperr.NewFromError(err, "Failed to fetch reviews by params").SetLocation()
	}

	query = p.conn.Rebind(query)


	rows, err := p.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, apperr.NewFromError(err, "Failed to fetch reviews by params").SetLocation()
	}
	defer rows.Close()

	var reviews []entity.ReviewWithRelations
	for rows.Next() {
		
		var review entity.ReviewWithRelations
		err := rows.Scan(
			&review.ID,
			&review.ProfessorID,
			&review.UserID,
			&review.Comment,
			&review.DiffRate,
			&review.FriendlyRate,
			&review.CreatedAt,
			&review.User.ID,
			&review.User.Username,
			&review.Professor.ID,
			&review.Professor.Name,
			&review.Professor.Faculty,
			&review.Professor.Major,
			&review.Professor.ProfileImgLink,
			&review.LikeCounter,
			&review.DislikeCounter,
		)
		if err != nil {
			return nil, apperr.NewFromError(err, "Failed to fetch reviews by params").SetLocation()
		}
		reviews = append(reviews, review)
	}

	return reviews, nil

}

func (p *reviewRepositoryImplPostgre) FetchRatingDistributionByProfId(ctx context.Context, profId string, column entity.RatingDistributionCol) (entity.RatingDistribution, error) {
	qb := goqu.From(goqu.T(reviewTableName).As("r")).Select(goqu.I("r.prof_id"))

	for i := 1; i <= 5; i++ {
		qb = qb.SelectAppend(
			goqu.COUNT(
				goqu.Case().
					When(goqu.L("FLOOR(?) = ?", goqu.I(column), i), 1).
					Else(nil),
			).As(fmt.Sprintf("rating_%d", i)),
		)
	}

	qb = qb.GroupBy(goqu.I("r.prof_id"))

	qb = qb.
		Where(goqu.I("r.prof_id").Eq(profId)).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return entity.RatingDistribution{}, apperr.NewFromError(err, "Failed to fetch rating distribution").SetLocation()
	}

	query = p.conn.Rebind(query)
	fmt.Println("QUERY: ", query)

	var res entity.RatingDistribution
	err = p.conn.QueryRowxContext(ctx, query, args...).StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.RatingDistribution{}, contracts.ErrProfessorNotFound
		}
		return entity.RatingDistribution{}, apperr.NewFromError(err, "Failed to fetch rating distribution").SetLocation()
	}

	return res, nil  

}

func (p *reviewRepositoryImplPostgre) GetReviewsItemsByParams(ctx context.Context, params *dto.FetchReviewParams) (uint64, error) {
	var count uint64
	
	qb := goqu.
		From(reviewTableName).
		Select(goqu.COUNT(goqu.Star()))

	if params.ProfId != "" {
		qb = qb.Where(goqu.I("reviews.prof_id").Eq(params.ProfId))
	}

	if params.UserId != "" {
		qb = qb.Where(goqu.I("reviews.user_id").Eq(params.UserId))
	}
		
	query, args, err := qb.SetDialect(goqu.GetDialect("postgres")).Prepared(true).ToSQL()
	if err != nil {
		return 0, apperr.NewFromError(err, "Failed to get reviews counter").SetLocation()
	}

	query = p.conn.Rebind(query)

	err = p.conn.QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, apperr.NewFromError(err, "Failed to get reviews counter").SetLocation()

	}

	return count, nil
}