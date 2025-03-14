package response

type ErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Err      error  `json:"error"`
	Location string `json:"-"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Meta    any    `json:"meta,omitempty"`
}

func NewResponse(
	code int,
	message string,
	data any,
	meta any,
) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) WithErr(err error) *ErrorResponse {
	e.Err = err
	return e
}

func (e *ErrorResponse) WithLocation(location string) *ErrorResponse {
	e.Location = location
	return e
}
