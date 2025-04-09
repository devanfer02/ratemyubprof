package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type ProfessorRepositoryProvider interface {
	NewClient(tx bool) (ProfessorRepository, error)
}

type ProfessorRepository interface {
	FetchAllProfessors(ctx context.Context, params *dto.FetchProfessorParam, pageQuery *dto.PaginationQuery) ([]entity.Professor, error)
	InsertProfessorsBulk(ctx context.Context, professors []entity.Professor) error 	
	InsertProfessorReview(ctx context.Context, review *entity.Review) error
}