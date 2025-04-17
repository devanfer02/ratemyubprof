package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID               string    `db:"id"`
	NIM              string    `db:"nim"`
	Username         string    `db:"username"`
	Password         string    `db:"password"`
	ForgotPasswordAt sql.NullTime `db:"forgot_password_at"`
	CreatedAt        time.Time `db:"created_at"`
}
