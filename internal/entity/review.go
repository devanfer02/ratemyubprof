package entity

import "time"

type RatingDistributionCol = string

var (
	DifficultyDistirbutionCol RatingDistributionCol = "difficulty_rating"
	FriendlyDistirbutionCol   RatingDistributionCol = "friendly_rating"
)

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
	User           User      `db:"user"`
	Professor      Professor `db:"professor"`
	IsLiked        int      `db:"is_liked"`
	LikeCounter    int       `db:"like_counter"`
	DislikeCounter int       `db:"dislike_counter"`
}

type RatingDistribution struct {
	ProfID  string `db:"prof_id" json:"-"`
	Rating1 int    `db:"rating_1" json:"ratingCounter1"`
	Rating2 int    `db:"rating_2" json:"ratingCounter2"`
	Rating3 int    `db:"rating_3" json:"ratingCounter3"`
	Rating4 int    `db:"rating_4" json:"ratingCounter4"`
	Rating5 int    `db:"rating_5" json:"ratingCounter5"`
}
