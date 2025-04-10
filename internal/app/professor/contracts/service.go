package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
)

type ProfessorService interface {
	FetchAllProfessors(ctx context.Context, params *dto.FetchProfessorParam, pageQuery *dto.PaginationQuery) ([]dto.FetchProfessorResponse, dto.PaginationResponse, error)
	FetchProfessorByID(ctx context.Context, id string) (dto.FetchProfessorResponse, error)
	
	FetchProfessorReviews(ctx context.Context, id string) ([]dto.FetchReviewResponse, error)
	CreateReview(ctx context.Context, param *dto.ProfessorReviewRequest) error
}