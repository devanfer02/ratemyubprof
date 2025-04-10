package repository

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/doug-martin/goqu/v9"
)

func (p *professorRepositoryImplPostgre) FetchProfessorReviews(ctx context.Context, id string, pageQuery *dto.PaginationQuery) ([]entity.ReviewWithRelations, error) {
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
			goqu.I("p.major").As(goqu.C("professor.profile_img_link")),
		).
		From(goqu.T(reviewTableName).As("r")).
		Join(
			goqu.T("users").As("u"),
			goqu.On(goqu.I("r.user_id").Eq(goqu.I("u.id"))),
		).
		Join(
			goqu.T("professors").As("p"),
			goqu.On(goqu.I("r.prof_id").Eq(goqu.I("p.id"))),
		). 
		Where(goqu.I("p.id").Eq(id))

	if pageQuery.Page != 0 && pageQuery.Limit != 0{
		qb = qb.Offset((pageQuery.Page - 1) * pageQuery.Limit).Limit(pageQuery.Limit)
	}

	query, args, err := qb.SetDialect(goqu.GetDialect("postgres")).Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}

	query = p.conn.Rebind(query)

	rows, err := p.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
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
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil

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