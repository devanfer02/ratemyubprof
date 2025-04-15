package dto

type FetchReviewResponse struct {
	ID           string                 `json:"id"`
	ProfessorID  string                 `json:"profId"`
	UserID       string                 `json:"userId"`
	Comment      string                 `json:"comment"`
	DiffRate     float32                `json:"difficultyRating"`
	FriendlyRate float32                `json:"friendlyRating"`
	CreatedAt    string                 `json:"createdAt"`
	Like         int                    `json:"like"`
	Dislike      int                    `json:"dislike"`
	User         FetchUserResponse      `json:"user"`
	Professor    FetchProfessorResponse `json:"professor"`
}

type FetchReviewParams struct {
	ProfId string `param:"profId"`
	UserId string `param:"userId"`
}
