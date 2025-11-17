package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/doug-martin/goqu/v9"
	"github.com/lib/pq"
)


func (u *userRepositoryImplPostgre) InsertUser(ctx context.Context, user *entity.User) error {
	qb := goqu.Insert(userTableName).Rows(
		goqu.Record{
			"id": user.ID,
			"nim": user.NIM,
			"username": user.Username,
			"password": user.Password,
		},
	)

	query, args, err := qb.SetDialect(goqu.GetDialect("postgres")).Prepared(true).ToSQL()
	if err != nil {
		return apperr.NewFromError(err, "Failed to insert user").SetLocation()
	}

	query = u.conn.Rebind(query)

	_, err = u.conn.ExecContext(ctx, query, args...)

	if err != nil {
		
		if contracts.IsErrorCode(err, contracts.PgsqlUniqueViolationErr) {
			pgErr, ok := err.(*pq.Error)
			if ok {
				if pgErr.Constraint == "users_nim_unique" {
					return contracts.ErrAlreadyRegistered
				} else if pgErr.Constraint == "users_username_unique" {
					return contracts.ErrUsernameTaken
				} 
			} 

			return contracts.ErrUsernameTaken
		}

		return apperr.NewFromError(err, "Failed to insert user").SetLocation()
	}

	return nil
}

func (u *userRepositoryImplPostgre) FetchUserByParams(ctx context.Context, params *dto.FetchUserParams) (entity.User, error) {
	qb := goqu.Select("*").
		From(userTableName)

	if params.Username != "" {
		qb = qb.Where(goqu.I("username").Eq(params.Username))
	}
	
	if params.NIM != "" {
		qb = qb.Where(goqu.I("nim").Eq(params.NIM))
	}

	if params.ID != "" {
		qb = qb.Where(goqu.I("id").Eq(params.ID))
	}

	query, args, err := qb.SetDialect(goqu.GetDialect("postgres")).ToSQL()
	if err != nil {
		return entity.User{}, apperr.NewFromError(err, "Failed to fetch user by username").SetLocation()
	}

	query = u.conn.Rebind(query)

	var user entity.User
	err = u.conn.QueryRowxContext(ctx, query, args...).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, contracts.ErrUserNotExists
		}

		return entity.User{}, apperr.NewFromError(err, "Failed to fetch user by username").SetLocation()
	}

	return user, nil
}


func (u *userRepositoryImplPostgre) UpdateUser(ctx context.Context, user *entity.User) error {
	qb := goqu.Update(userTableName).
		Set(goqu.Record{
			"password": user.Password,
			"forgot_password_at": user.ForgotPasswordAt,
		}). 
		Where(goqu.C("nim").Eq(user.NIM)). 
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()

	if err != nil {
		return apperr.NewFromError(err, "Failed to update user").SetLocation()
	}

	query = u.conn.Rebind(query)	

	res, err := u.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return apperr.NewFromError(err, "Failed to update user").SetLocation()
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return apperr.NewFromError(err, "Failed to update user").SetLocation()
	}

	if rows == 0 {
		return contracts.ErrUserNotExists
	}

	if rows > 1 {
		return contracts.ErrWeirdUpdate
	}
	
	return nil 
}

func (u *userRepositoryImplPostgre) FetchUserProfile(ctx context.Context, userID string) (dto.UserProfileResponse, error) {
	var userProfile dto.UserProfileResponse

	// Fetch user
	userQb := goqu.Select(
		goqu.I("u.id"),
		goqu.I("u.nim"),
		goqu.I("u.username"),
		goqu.I("u.created_at"),
	).From(goqu.T("users").As("u")).Where(goqu.I("u.id").Eq(userID))

	query, args, err := userQb.SetDialect(goqu.GetDialect("postgres")).ToSQL()
	if err != nil {
		return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to fetch user profile").SetLocation()
	}

	query = u.conn.Rebind(query)

	err = u.conn.QueryRowxContext(ctx, query, args...).StructScan(&userProfile)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.UserProfileResponse{}, contracts.ErrUserNotExists
		}
		return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to fetch user profile").SetLocation()
	}

	// Fetch reviews count
	countQb := goqu.Select(goqu.COUNT("*")).From("reviews").Where(goqu.I("user_id").Eq(userID))
	query, args, err = countQb.SetDialect(goqu.GetDialect("postgres")).ToSQL()
	if err != nil {
		return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to fetch user profile").SetLocation()
	}

	query = u.conn.Rebind(query)

	err = u.conn.QueryRowxContext(ctx, query, args...).Scan(&userProfile.ReviewsCount)
	if err != nil {
		return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to fetch user profile").SetLocation()
	}

	// Fetch recent reviews
	reviewsQb := goqu.Select(
		goqu.I("r.id"),
		goqu.I("r.prof_id"),
		goqu.I("r.user_id"),
		goqu.I("r.comment"),
		goqu.I("r.difficulty_rating"),
		goqu.I("r.friendly_rating"),
		goqu.I("r.created_at"),
		goqu.L("jsonb_build_object").As("professor"),
	).From(goqu.T("reviews").As("r")).
		LeftJoin(goqu.T("professors").As("p"), goqu.On(goqu.I("r.prof_id").Eq(goqu.I("p.id")))).
		Where(goqu.I("r.user_id").Eq(userID)).
		Order(goqu.I("r.created_at").Desc()).
		Limit(5)

	query, args, err = reviewsQb.SetDialect(goqu.GetDialect("postgres")).ToSQL()
	if err != nil {
		return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to fetch user profile").SetLocation()
	}

	query = u.conn.Rebind(query)

	rows, err := u.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to fetch user profile").SetLocation()
	}

	defer rows.Close()

	for rows.Next() {
		var review dto.FetchReviewResponse
		var professorJSON []byte
		err := rows.Scan(
			&review.ID,
			&review.ProfessorID,
			&review.UserID,
			&review.Comment,
			&review.DiffRate,
			&review.FriendlyRate,
			&review.CreatedAt,
			&professorJSON,
		)
		if err != nil {
			return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to scan review").SetLocation()
		}
		
		if err := json.Unmarshal(professorJSON, &review.Professor); err != nil {
			return dto.UserProfileResponse{}, apperr.NewFromError(err, "Failed to unmarshal professor data").SetLocation()
		}

		userProfile.RecentReviews = append(userProfile.RecentReviews, review)
	}

	return userProfile, nil
}