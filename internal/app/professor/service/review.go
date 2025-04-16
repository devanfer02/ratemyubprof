package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
	"github.com/oklog/ulid/v2"
)

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

func (s *professorService) UpdateProfessorReview(ctx context.Context, req *dto.ProfessorReviewRequest) error {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return err 
	}

	entity := formatter.FormatReviewToEntity(req)
	err = repoClient.UpdateProfessorReview(ctx, &entity)
	if err != nil {
		return err 
	}

	return nil 
}

func (s *professorService) DeleteProfessorReview(ctx context.Context, params *dto.FetchReviewParams) error {
	repoClient, err := s.profRepo.NewClient(false)
	if err != nil {
		return err 
	}

	err = repoClient.DeleteProfessorReview(ctx, params)
	if err != nil {
		return err 
	}

	return nil 
}