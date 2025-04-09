package service

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/bytedance/sonic"
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

func (s *professorService) FetchStaticProfessorData(param *dto.FetchProfessorParam) ([]dto.ProfessorStatic, error) {
	var (
		err error 
		professors []dto.ProfessorStatic
		fileName = "data/dosenub.json"
	)

	file, err := os.Open(fileName)
	if err != nil {
		s.logger.Error("[ProfessorService.FetchStaticProfessorData] failed to open file", zap.Error(err))
		return nil, err 
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		s.logger.Error("[ProfessorService.FetchStaticProfessorData] failed to read file", zap.Error(err))
		return nil, err
	}

	if err := sonic.Unmarshal(data, &professors); err != nil {
		s.logger.Error("[ProfessorService.FetchStaticProfessorData] failed to unmarshal data", zap.Error(err))
		return nil, err
	}

	professors = util.Filter(professors, func(p dto.ProfessorStatic) bool {
		if param.Name != "" && !strings.Contains(strings.ToLower(p.Name), strings.ToLower(param.Name)) {
			return false
		}
		if param.Faculty != "" && !strings.Contains(strings.ToLower(p.Fakultas), strings.ToLower(param.Faculty)) {
			return false
		}
		if param.Major != "" && !strings.Contains(strings.ToLower(p.Prodi), strings.ToLower(param.Major)) {
			return false
		}
		return true
	})

	return professors, nil
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