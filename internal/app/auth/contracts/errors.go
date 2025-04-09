package contracts

import (
	"net/http"

	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
)

var (
	ErrRequestTimeout    = apperr.New(http.StatusRequestTimeout, "http request timeout")
	ErrInvalidCredential = apperr.New(http.StatusUnauthorized, "invalid credential")
	ErrInvalidToken      = apperr.New(http.StatusUnauthorized, "invalid token")
)