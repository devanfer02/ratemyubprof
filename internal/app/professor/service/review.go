package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
	"github.com/oklog/ulid/v2"
)

func (s *professorService) FetchProfessorReviews(ctx context.Context, id string, pageQuery *dto.PaginationQuery) ([]dto.FetchReviewResponse, dto.PaginationResponse,error) {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return nil, dto.PaginationResponse{},err 
	}

	entities, err := repoClient.FetchProfessorReviews(ctx, id, pageQuery)
	if err != nil {
		return nil, dto.PaginationResponse{},err 
	}

	items, err := repoClient.GetReviewsItemsByProfID(ctx, id)
	if err != nil {
		return nil, dto.PaginationResponse{},err 
	}

	pageMeta := util.GetPagination(uint(items), pageQuery.Limit, pageQuery.Page)

	res := formatter.FormatReviewEntitiesToDto(entities)

	return res, pageMeta,nil 
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