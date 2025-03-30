package apperr

type AppError struct {
	Code     int
	Message  string
	Location string
	Err      error
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) WithErr(err error) *AppError {
	e.Err = err
	return e
}

func (e *AppError) WithLocation(location string) *AppError {
	e.Location = location
	return e
}

