package response

import "errors"

type Response struct {
	Message  string `json:"message"`
	Data     any    `json:"data"`
	Meta     any    `json:"meta,omitempty"`
	Err      error  `json:"error,omitempty"`
	Location string `json:"-"`
	Code     int    `json:"-"`
}

func New(
	message string,
	data any,
	meta any,
) *Response {
	return &Response{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func NewErr(
	code int,
	message string,
) *Response {
	return &Response{
		Message: message,
		Code: code,
		Err: errors.New(message),
	}
}

func (r Response) Error() string {
	return r.Message
}

func (r *Response) WithErr(err error) *Response {
	r.Err = err
	return r
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code 
	return r 
}

func (r *Response) WithLocation(location string) *Response {
	r.Location = location
	return r
}
