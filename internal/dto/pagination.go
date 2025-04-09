package dto

type PaginationResponse struct {
	TotalItems uint `json:"totalItems"`
	TotalPages uint `json:"totalPages"`
	Current    uint `json:"current"`
	Next       uint `json:"next"`
	Prev       uint `json:"prev"`
}

type PaginationQuery struct {
	Page  uint `query:"page" validate:"min=1"`	
	Limit uint `query:"limit" validate:"max=100"`
}