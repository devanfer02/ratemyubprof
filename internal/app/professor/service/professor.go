package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
)


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

	items, err := repoClient.GetProfessorItems(ctx, params)
	if err != nil {
		return nil, dto.PaginationResponse{}, err 
	}

	pageMeta := util.GetPagination(uint(items), pageQuery.Limit, pageQuery.Page)

	responses := formatter.FormatProfessorEntitiesToDto(professors)

	return responses, pageMeta, nil
}

func (s *professorService) FetchProfessorByID(ctx context.Context, id string) (dto.ProfessorResponse, error) {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return dto.ProfessorResponse{}, err 
	}

	professor, err := repoClient.FetchProfessorByID(ctx, id)
	if err != nil {
		return dto.ProfessorResponse{}, err 
	}

	response := formatter.FormatProfessorEntityToDto(professor)

	return response, nil 
}

