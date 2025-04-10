package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
)

type professorService struct {
	profRepo contracts.ProfessorRepositoryProvider
}

func NewProfessorService(profRepo contracts.ProfessorRepositoryProvider) contracts.ProfessorService {
	return &professorService{
		profRepo: profRepo,
	}
}
