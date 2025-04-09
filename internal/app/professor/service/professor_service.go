package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/util"
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

func (s *professorService) FetchAllProfessors(
	ctx context.Context, 
	params *dto.FetchProfessorParam, 
	pageQuery *dto.PaginationQuery,
) (
	[]dto.ProfessorResponse, 
	dto.PaginationResponse,
	error,
) {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return nil, dto.PaginationResponse{}, err 
	}

	professors, err := repoClient.FetchAllProfessors(ctx, params, pageQuery)
	if err != nil {
		return nil, dto.PaginationResponse{}, err 
	}

	items, err := repoClient.CountProfessor(ctx)
	if err != nil {
		return nil, dto.PaginationResponse{}, err 
	}

	pageMeta := util.GetPagination(uint(items), pageQuery.Limit, pageQuery.Page)

	responses := formatter.FormatProfessorEntityToDto(professors)

	return responses, pageMeta, nil
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