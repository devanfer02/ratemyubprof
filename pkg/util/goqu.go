package util

import (
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/doug-martin/goqu/v9"
)

func AddParamsToFetchProf(qb *goqu.SelectDataset, params *dto.FetchProfessorParam) *goqu.SelectDataset {
	if params.Faculty != "" {
		qb = qb.Where(goqu.C("faculty").ILike("%" + params.Faculty + "%"))
	}

	if params.Major != "" {
		qb = qb.Where(goqu.C("major").ILike("%" + params.Major + "%"))
	}

	if params.Name != "" {
		qb = qb.Where(goqu.C("name").ILike("%" + params.Name + "%"))
	}

	return qb
}