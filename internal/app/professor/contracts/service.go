package contracts

import "github.com/devanfer02/ratemyubprof/internal/dto"

type ProfessorService interface {
	FetchStaticProfessorData(param *dto.FetchProfessorParam) ([]dto.ProfessorStatic, error) 
}