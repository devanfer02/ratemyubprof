package dto

type FetchReviewResponse struct {
	ID           string                 `json:"id"`
	ProfessorID  string                 `json:"prof_id"`
	UserID       string                 `json:"user_id"`
	Comment      string                 `json:"comment"`
	DiffRate     float32                `json:"difficulty_rating"`
	FriendlyRate float32                `json:"friendly_rating"`
	CreatedAt    string                 `json:"created_at"`
	User         FetchUserResponse      `json:"user"`
	Professor    FetchProfessorResponse `json:"professor"`
}
