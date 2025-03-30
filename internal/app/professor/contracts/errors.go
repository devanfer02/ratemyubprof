package contracts

import (
	"net/http"

	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
)

var (
	ErrRequestTimeout = apperr.New(http.StatusRequestTimeout, "request timeout")
)