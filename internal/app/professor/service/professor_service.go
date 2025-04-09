package service

import (
	"context"
	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
	"github.com/oklog/ulid/v2"
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

func (s *professorService) FetchAllProfessors(ctx context.Context, params *dto.FetchProfessorParam, pageQuery *dto.PaginationQuery) ([]dto.ProfessorResponse, error) {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return nil, err 
	}

	professors, err := repoClient.FetchAllProfessors(ctx, params, pageQuery)
	if err != nil {
		return nil, err 
	}

	responses := formatter.FormatProfessorEntityToDto(professors)

	return responses, nil
}

func (s *professorService) CreateReview(ctx context.Context, param *dto.ProfessorReviewRequest) error {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return err 
	}

	entity := formatter.FormatReviewToEntity(param)
	entity.ID = ulid.Make().String()
	err = repoClient.InsertProfessorReview(ctx, &entity)

	if err != nil {
		return err 
	}


	return nil
}