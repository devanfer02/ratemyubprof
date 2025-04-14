package dto

type ReviewReactionRequest struct {
	ReviewID string `param:"id" validate:"required"`
	UserID   string 
	Type     string   `json:"type" validate:"required,reactionType"`
}