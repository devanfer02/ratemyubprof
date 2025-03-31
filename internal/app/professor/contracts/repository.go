package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type ProfessorRepositoryProvider interface {
	NewClient(tx bool) (ProfessorRepository, error)
}

type ProfessorRepository interface {
	InsertProfessorsBulk(ctx context.Context, professors []entity.Professor) error 	
}