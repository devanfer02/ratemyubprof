package dto

type ProfessorStatic struct {
	Name     string `json:"name"`
	Fakultas string `json:"fakultas"`
	Prodi    string `json:"prodi"`
	ImgLink  string `json:"img"`
}

type ProfessorReviewRequest struct {
	ProfessorID  string  `param:"id" validate:"required"`
	UserID       string  
	Comment      string  `json:"comment" validate:"required"`
	DiffRate     float32 `json:"diffRate" validate:"required,min=1,max=5"`
	FriendlyRate float32 `json:"friendlyRate" validate:"required,min=1,max=5"`
}

type FetchProfessorParam struct {
	Name    string
	Faculty string
	Prodi   string
}
