package service

import (
	"context"
	"sync"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
)


func (s *professorService) FetchAllProfessors(
	ctx context.Context, 
	params *dto.FetchProfessorParam, 
	pageQuery *dto.PaginationQuery,
) (
	[]dto.FetchProfessorResponse, 
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

func (s *professorService) FetchProfessorByID(ctx context.Context, id string) (dto.FetchProfessorResponse, dto.ProfessorRatingDistribution,error) {
	profRepo, err := s.profRepo.NewClient(false)
	if err != nil {
		return dto.FetchProfessorResponse{}, dto.ProfessorRatingDistribution{}, err 
	}

	reviewRepo, err := s.reviewRepo.NewClient(false)
	if err != nil {
		return dto.FetchProfessorResponse{}, dto.ProfessorRatingDistribution{}, err
	}

	var (
		wg sync.WaitGroup
		errChan = make(chan error, 3)
		difRate = make(chan entity.RatingDistribution, 1)
		frdRate = make(chan entity.RatingDistribution, 1)
		profRes = make(chan entity.Professor, 1)

		distribution = dto.ProfessorRatingDistribution{}

		actions = []func(){
			func() {
				defer wg.Done()
				dif, err := reviewRepo.FetchRatingDistributionByProfId(ctx, id, entity.DifficultyDistirbutionCol)
				
				if err != nil {
					errChan <- err
					return 
				}

				difRate <- dif
			},
			func() {
				defer wg.Done()
				frd, err := reviewRepo.FetchRatingDistributionByProfId(ctx, id, entity.FriendlyDistirbutionCol)
				
				if err != nil {
					errChan <- err
					return 
				}

				frdRate <- frd
			},
			func() {
				defer wg.Done()
				professor, err := profRepo.FetchProfessorByID(ctx, id)
				if err != nil {
					errChan <- err
					return 
				}

				profRes <- professor
			},
		}
	)


	for _, action := range actions {
		wg.Add(1)
		go action()
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(difRate)
		close(frdRate)
		close(profRes)
	}()

	for err := range errChan {
		if err != nil {
			return dto.FetchProfessorResponse{}, dto.ProfessorRatingDistribution{}, err 
		}
	}

	professor := <- profRes  
	distribution.ProfessorID = professor.ID
	distribution.DiffcultyDistribution = <- difRate
	distribution.FriendlyDistirbutuion = <- frdRate

	response := formatter.FormatProfessorEntityToDto(professor)

	return response, distribution, nil 
}

