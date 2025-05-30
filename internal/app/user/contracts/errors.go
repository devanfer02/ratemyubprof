package contracts

import (
	"net/http"

	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/lib/pq"
)

var (
	ErrRequestTimeout    = apperr.New(http.StatusRequestTimeout, "http request timeout")
	ErrUsernameTaken     = apperr.New(http.StatusConflict, "username is already taken")
	ErrAlreadyRegistered = apperr.New(http.StatusConflict, "student already registered")
	ErrUserNotExists     = apperr.New(http.StatusNotFound, "user does not exist")
	ErrWeirdUpdate       = apperr.New(http.StatusInternalServerError, "weird behaviour update user")
	ErrResetPasswordLimit = apperr.New(http.StatusTooManyRequests, "reset password limit reached")

	PgsqlUniqueViolationErr = pq.ErrorCode("23505")
)

func IsErrorCode(err error, code pq.ErrorCode) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == code
	}

	return false
}
