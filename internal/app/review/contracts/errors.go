package contracts

import (
	"net/http"

	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
)

var (
	ErrProfessorNotFound = apperr.New(http.StatusNotFound, "professor not found")
)