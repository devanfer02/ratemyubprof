package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
	"github.com/oklog/ulid/v2"
)

func (s *professorService) FetchProfessorReviews(ctx context.Context, id string) ([]dto.FetchReviewResponse, error) {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return nil, err 
	}

	entities, err := repoClient.FetchProfessorReviews(ctx, id)
	if err != nil {
		return nil, err 
	}

	res := formatter.FormatReviewEntitiesToDto(entities)

	return res, nil 
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