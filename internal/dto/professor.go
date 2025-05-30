package dto

type ProfessorStatic struct {
	Name     string `json:"name"`
	Fakultas string `json:"fakultas"`
	Prodi    string `json:"prodi"`
	ImgLink  string `json:"img"`
}

type ProfessorReviewRequest struct {
	ProfessorID  string `param:"id" validate:"required"`
	UserID       string
	Comment      string  `json:"comment" validate:"required"`
	DiffRate     float32 `json:"diffRate" validate:"required,min=1,max=5"`
	FriendlyRate float32 `json:"friendlyRate" validate:"required,min=1,max=5"`
}

type FetchProfessorResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Faculty         string  `json:"faculty"`
	Major           string  `json:"major"`
	ProfileImgLink  string  `json:"profileImgLink"`
	ReviewsCount    uint64  `json:"reviewsCount"`
	AvgDiffRate     float32 `json:"avgDiffRate"`
	AvgFriendlyRate float32 `json:"avgFriendlyRate"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

type FetchProfessorParam struct {
	Name    string `query:"name"`
	Faculty string `query:"faculty"`
	Major   string `query:"major"`
}
