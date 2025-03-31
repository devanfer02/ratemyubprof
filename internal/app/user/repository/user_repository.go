package repository

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

var (
	userTableName = "users"
)

type repository struct {
	conn *sqlx.DB
}

type userRepositoryImplPostgre struct {
	conn sqlx.ExtContext
}

func NewUserRepository(conn *sqlx.DB) contracts.UserRepositoryProvider {
	return &repository{
		conn: conn,
	}
}

func (r *repository) NewClient(tx bool) (contracts.UserRepository, error) {
	var (
		conn  sqlx.ExtContext
		err error
	)

	if tx {
		conn, err = r.conn.Beginx()
		if err != nil {
			return nil, err
		}
	} else {
		conn = r.conn
	}

	return &userRepositoryImplPostgre{
		conn: conn,
	}, nil
}


func (u *userRepositoryImplPostgre) InsertUser(ctx context.Context, user *entity.User) error {
	qb := goqu.Insert(userTableName).Rows(
		goqu.Record{
			"id": user.ID,
			"username": user.Username,
			"password": user.Password,
		},
	)

	query, args, err := qb.SetDialect(goqu.GetDialect("postgres")).Prepared(true).ToSQL()
	if err != nil {
		return err
	}

	query = u.conn.Rebind(query)

	_, err = u.conn.ExecContext(ctx, query, args...)
	if err != nil {
		
		if contracts.IsErrorCode(err, contracts.PgsqlUniqueViolationErr) {
			return contracts.ErrUsernameTaken
		}

		return err
	}

	return nil
}

func (u *userRepositoryImplPostgre) FetchUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	qb := goqu.Select("*").
		From(userTableName).
		Where(goqu.C("username").Eq(username))

	query, args, err := qb.SetDialect(goqu.GetDialect("postgres")).ToSQL()
	if err != nil {
		return nil, err
	}

	query = u.conn.Rebind(query)

	var user entity.User
	rows, err := u.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
