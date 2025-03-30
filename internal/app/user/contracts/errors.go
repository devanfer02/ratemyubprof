package contracts

import (
	"net/http"

	"github.com/devanfer02/ratemyubprof/pkg/response"
)

var (
	ErrRequestTimeout = response.NewErr(http.StatusRequestTimeout, "http request timeout")
)