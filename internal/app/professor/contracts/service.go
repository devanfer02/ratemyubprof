package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
)

type ProfessorService interface {
	FetchStaticProfessorData(param *dto.FetchProfessorParam) ([]dto.ProfessorStatic, error) 
	CreateReview(ctx context.Context, param *dto.ProfessorReviewRequest) error
}