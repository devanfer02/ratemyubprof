package entity

type ReviewReaction struct {
	UserID    string `db:"user_id"`
	ReviewID  string `db:"review_id"`
	Type      uint   `db:"reactino_type"`
	CreatedAt string `db:"created_at"`
}
