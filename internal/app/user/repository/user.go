package repository

import (
	"context"
	"database/sql"

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
		}). 
		Where(goqu.I("nim").Eq(user.NIM)). 
		Prepared(true). 
		SetDialect(goqu.GetDialect("postgres"))

	query, args, err := qb.ToSQL()

	if err != nil {
		return apperr.NewFromError(err, "Failed to update user").SetLocation()
	}

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