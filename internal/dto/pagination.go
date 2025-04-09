package dto

type PaginationResponse struct {
	TotalItems uint64 `json:"totalItems"`
	TotalPages uint32 `json:"totalPages"`
	Current    uint32 `json:"current"`
	Next       uint32 `json:"next"`
	Prev       uint32 `json:"prev"`
}

type PaginationQuery struct {
	Page  uint `query:"page" validate:"min=1"`	
	Limit uint `query:"limit" validate:"max=100"`
}