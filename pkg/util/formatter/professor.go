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

func FormatProfessorEntitiesToDto(professors []entity.Professor) []dto.FetchProfessorResponse {
	var professorsResponse []dto.FetchProfessorResponse
	for _, professor := range professors {
		professorsResponse = append(professorsResponse, FormatProfessorEntityToDto(professor))
	}
	return professorsResponse
}

func FormatProfessorEntityToDto(professor entity.Professor) dto.FetchProfessorResponse {
	return dto.FetchProfessorResponse{
		ID:             professor.ID,
		Name:           professor.Name,
		Faculty:        professor.Faculty,
		Major:          professor.Major,
		ProfileImgLink: professor.ProfileImgLink,
		CreatedAt:      professor.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      professor.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}