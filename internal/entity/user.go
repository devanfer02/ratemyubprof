package entity

import "time"

type User struct {
	ID        string    `db:"id"`
	NIM       string    `db:"nim"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
