package dto

import "github.com/devanfer02/ratemyubprof/internal/entity"

type FetchReviewResponse struct {
	ID           string                 `json:"id"`
	ProfessorID  string                 `json:"profId"`
	UserID       string                 `json:"userId"`
	Comment      string                 `json:"comment"`
	DiffRate     float32                `json:"difficultyRating"`
	FriendlyRate float32                `json:"friendlyRating"`
	CreatedAt    string                 `json:"createdAt"`
	IsLiked      int                    `json:"isLiked"`
	Like         int                    `json:"like"`
	Dislike      int                    `json:"dislike"`
	User         FetchUserResponse      `json:"user"`
	Professor    FetchProfessorResponse `json:"professor"`
}

type FetchReviewParams struct {
	ID         string  
	ProfId     string `param:"profId"`
	UserId     string `param:"userId"`
	SignedUser string
}

type ProfessorRatingDistribution struct {
	ProfessorID           string                    `json:"profId"`
	DiffcultyDistribution entity.RatingDistribution `json:"difficultyDistribution"`
	FriendlyDistirbutuion entity.RatingDistribution `json:"friendlyDistribution"`
}
