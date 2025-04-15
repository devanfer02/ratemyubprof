package tests

import (
	"testing"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/app/user/repository"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/devanfer02/ratemyubprof/tests/app/fixtures"
	"github.com/jmoiron/sqlx"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	type args struct {
		user entity.User
	}

	tests := []struct {
		name       string
		wantErr    bool
		wantErrMsg error
		args       args
		beforeTest func(
			args args,
			db *sqlx.DB,
		) error
	}{
		{
			name:    "When inserting user, it should return success",
			wantErr: false,
			args: args{
				user: entity.User{
					NIM:      "225150200111",
					Username: "user",
					Password: "password",
				},
			},
		},
		{
			name:       "When inserting duplicate user's username, it should return error username taken",
			wantErr:    true,
			wantErrMsg: contracts.ErrUsernameTaken,
			args: args{
				user: entity.User{
					ID:       ulid.Make().String(),
					NIM:      "225150200111",
					Username: "user",
					Password: "password",
				},
			},
			beforeTest: func(args args, db *sqlx.DB) error {
				args.user.ID = ulid.Make().String()
				args.user.NIM = "should be unique"
				query := `INSERT INTO users (id, nim, username, password) VALUES (:id, :nim, :username, :password)`
				_, err := db.NamedExec(query, args.user)
				return err
			},
		},
		{
			name:       "When inserting user with existing NIM, it should return error user already registered",
			wantErr:    true,
			wantErrMsg: contracts.ErrAlreadyRegistered,
			args: args{
				user: entity.User{
					ID:       ulid.Make().String(),
					NIM:      "225150200111",
					Username: "user",
					Password: "password",
				},
			},
			beforeTest: func(args args, db *sqlx.DB) error {
				args.user.ID = ulid.Make().String()
				args.user.Username = "should be unique"
				query := `INSERT INTO users (id, nim, username, password) VALUES (:id, :nim, :username, :password)`
				_, err := db.NamedExec(query, args.user)
				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dbx, clean := fixtures.NewDB()
			defer clean()

			repo := repository.NewUserRepository(dbx)
			client, err := repo.NewClient(false)

			if err != nil {
				t.Fatalf("Failed to initalize client | ERR: %v", err.Error())
			}

			if test.beforeTest != nil {
				err := test.beforeTest(test.args, dbx)

				if err != nil {
					t.Fatalf("Failed to prep test with beforeTest func | ERR: %v", err.Error())
				}
			}

			err = client.InsertUser(t.Context(), &test.args.user)

			if test.wantErr {
				assert.NotNil(t, err, "Expecting error to be thrown")

				if err != nil {
					assert.Equal(t, test.wantErrMsg, err, "Expecting same error result")
				}

			} else {
				assert.Nil(t, err, "Error should not be expected")
			}
		})
	}
}

func TestFetchUserByUsername(t *testing.T) {
	tests := []struct {
		name       string
		wantErr    error 
		want       *entity.User
		beforeTest func(
			user *entity.User,
			db *sqlx.DB,
		) error
	}{
		{
			name:    "When fetching user with existing username, it should return no error",
			wantErr: nil,
			want: &entity.User{
				ID: "myid",
				NIM: "nim",
				Username: "dia",
				Password: "passwordwkwk",
			},
			beforeTest: func(user *entity.User, db *sqlx.DB) error {
				query := `INSERT INTO users (id, nim, username, password) VALUES (:id, :nim, :username, :password)`
				_, err := db.NamedExec(query, user)
				return err
			},
		},
		{
			name:    "When fetching user with non-existing username, it should return error item does not exists",
			wantErr: contracts.ErrUserNotExists,
			want: &entity.User{
				Username: "dia",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dbx, clean := fixtures.NewDB()
			defer clean()

			repo := repository.NewUserRepository(dbx)
			client, err := repo.NewClient(false)

			if err != nil {
				t.Fatalf("Failed to initalize client | ERR: %v", err.Error())
			}

			if test.beforeTest != nil {
				err := test.beforeTest(test.want, dbx)

				if err != nil {
					t.Fatalf("Failed to prep test with beforeTest func | ERR: %v", err.Error())
				}
			}

			user, err := client.FetchUserByUsername(t.Context(), test.want.Username)

			if test.wantErr != nil {
				assert.NotNil(t, err, "Expecting error to be thrown")

				if err != nil {
					assert.Equal(t, test.wantErr, err, "Expecting same error result")
				}

			} else {
				assert.Nil(t, err, "Error should not be expected")
				assert.Equal(t, test.want.ID, user.ID, "Expecting same user result")
				assert.Equal(t, test.want.NIM, user.NIM, "Expecting same user result")
				assert.Equal(t, test.want.Username, user.Username, "Expecting same user result")
				assert.Equal(t, test.want.Password, user.Password, "Expecting same user result")
			}
		})
	}
}
