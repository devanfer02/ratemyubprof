package dto

type ReviewReactionRequest struct {
	ReviewID string `json:"review_id" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Type     string   `json:"type" validate:"required"`
}