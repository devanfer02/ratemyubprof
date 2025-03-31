package formatter

import (
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/oklog/ulid/v2"
)

func FormatProfessorStaticToEntity(professors []dto.ProfessorStatic) []entity.Professor {
	var professorsEntity []entity.Professor
	for _, professor := range professors {
		professorsEntity = append(professorsEntity, entity.Professor{
			ID: ulid.Make().String(),
			Name:    professor.Name,
			Faculty: professor.Fakultas,
			Major:   professor.Prodi,
			ProfileImgLink: professor.ImgLink,
		})
	}
	return professorsEntity

}