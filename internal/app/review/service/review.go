package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
)

func (s *reviewService) FetchReviewsByParams(ctx context.Context, params *dto.FetchReviewParams, pageQuery *dto.PaginationQuery) ([]dto.FetchReviewResponse, dto.PaginationResponse,error) {
	repoClient, err := s.reviewRepo.NewClient(false)
	if err != nil {
		return nil, dto.PaginationResponse{},err 
	}

	entities, err := repoClient.FetchReviewsByParams(ctx, params, pageQuery)
	if err != nil {
		return nil, dto.PaginationResponse{},err 
	}

	items, err := repoClient.GetReviewsItemsByParams(ctx, params)
	if err != nil {
		return nil, dto.PaginationResponse{},err 
	}

	pageMeta := util.GetPagination(uint(items), pageQuery.Limit, pageQuery.Page)

	res := formatter.FormatReviewEntitiesToDto(entities)

	return res, pageMeta,nil 
}