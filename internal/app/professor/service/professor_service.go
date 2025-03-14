package service

import (
	"io"
	"os"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/devanfer02/presentia-api/internal/app/professor/contracts"
	"github.com/devanfer02/presentia-api/internal/dto"
	"github.com/devanfer02/presentia-api/pkg/helpers"
)

type professorService struct {

}

func NewProfessorService() contracts.ProfessorService {
	return &professorService{}
}

func (s *professorService) FetchStaticProfessorData(param *dto.FetchProfessorParam) ([]dto.ProfessorStatic, error) {
	var (
		err error 
		professors []dto.ProfessorStatic
		fileName = "data/dosenub.json"
	)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err 
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := sonic.Unmarshal(data, &professors); err != nil {
		return nil, err
	}

	professors = helpers.Filter(professors, func(p dto.ProfessorStatic) bool {
		if param.Name != "" && !strings.Contains(strings.ToLower(p.Name), strings.ToLower(param.Name)) {
			return false
		}
		if param.Faculty != "" && !strings.Contains(strings.ToLower(p.Fakultas), strings.ToLower(param.Faculty)) {
			return false
		}
		if param.Prodi != "" && !strings.Contains(strings.ToLower(p.Prodi), strings.ToLower(param.Prodi)) {
			return false
		}
		return true
	})

	return professors, nil
}