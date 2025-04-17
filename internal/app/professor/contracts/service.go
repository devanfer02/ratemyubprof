package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
)

type ProfessorService interface {
	FetchAllProfessors(ctx context.Context, params *dto.FetchProfessorParam, pageQuery *dto.PaginationQuery) ([]dto.FetchProfessorResponse, dto.PaginationResponse, error)
	FetchProfessorByID(ctx context.Context, id string) (dto.FetchProfessorResponse, dto.ProfessorRatingDistribution,error)

	CreateReview(ctx context.Context, req *dto.ProfessorReviewRequest) error
	UpdateProfessorReview(ctx context.Context, req *dto.ProfessorReviewRequest) error 
	DeleteProfessorReview(ctx context.Context, params *dto.FetchReviewParams) error
}