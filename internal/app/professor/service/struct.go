package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"go.uber.org/zap"
)

type professorService struct {
	profRepo contracts.ProfessorRepositoryProvider
	logger *zap.Logger
}

func NewProfessorService(logger *zap.Logger, profRepo contracts.ProfessorRepositoryProvider) contracts.ProfessorService {
	return &professorService{
		logger: logger,
		profRepo: profRepo,
	}
}
