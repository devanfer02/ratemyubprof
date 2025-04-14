package contracts

import (
	"net/http"

	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/lib/pq"
)

var (
	ErrRequestTimeout      = apperr.New(http.StatusRequestTimeout, "request timeout")
	ErrItemNotFound        = apperr.New(http.StatusNotFound, "item not found")
	ErrMoreThanOneAffected = apperr.New(http.StatusNotFound, "affected item more than one")
	ErrItemAlreadyExists   = apperr.New(http.StatusConflict, "item already exists")

	PgsqlUniqueViolationErr = pq.ErrorCode("23505")	
)

func IsErrorCode(err error, code pq.ErrorCode) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == code
	}

	return false
}
