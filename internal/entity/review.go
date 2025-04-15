package entity

import "time"

type Review struct {
	ID           string    `db:"id"`
	ProfessorID  string    `db:"prof_id"`
	UserID       string    `db:"user_id"`
	Comment      string    `db:"comment"`
	DiffRate     float32   `db:"difficulty_rating"`
	FriendlyRate float32   `db:"friendly_rating"`
	CreatedAt    time.Time `db:"created_at"`
}

type ReviewWithRelations struct {
	Review
	User        User      `db:"user"`
	Professor   Professor `db:"professor"`
	LikeCounter int       `db:"like_counter"`
	DislikeCounter int       `db:"dislike_counter"`
}
