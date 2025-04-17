package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	review "github.com/devanfer02/ratemyubprof/internal/app/review/contracts"
)

type professorService struct {
	profRepo contracts.ProfessorRepositoryProvider
	reviewRepo  review.ReviewRepositoryProvider
}

func NewProfessorService(profRepo contracts.ProfessorRepositoryProvider, reviewRepo review.ReviewRepositoryProvider) contracts.ProfessorService {
	return &professorService{
		profRepo: profRepo,
		reviewRepo: reviewRepo,
	}
}
