package dto

type ProfessorStatic struct {
	Name     string `json:"name"`
	Fakultas string `json:"fakultas"`
	Prodi    string `json:"prodi"`
	ImgLink  string `json:"img"`
}

type FetchProfessorParam struct {
	Name    string
	Faculty string
	Prodi   string
}
