package apperr

import "runtime"

type AppError struct {
	Code    int
	Message string
	File    string
	Action  string
	Line    int
	Err     error
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func NewFromError(err error, message string) *AppError {
	return &AppError{
		Code:    500,
		Message: message,
		Err:     err,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) WithErr(err error) *AppError {
	e.Err = err
	return e
}

func (e *AppError) WithLocation(file string, method string, line int) *AppError {
	e.Action = method
	e.File = file
	e.Line = line 
	return e
}

func (e *AppError) SetLocation() *AppError {
	pc, file, line, _ := runtime.Caller(1)
	method := runtime.FuncForPC(pc).Name()

	e.Action = method 
	e.File = file 
	e.Line = line

	return e 
}

func (e *AppError) SetCode(code int) *AppError {
	e.Code = code
	return e
}