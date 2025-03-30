package response

type Response struct {
	Code    int    `json:"int"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Meta    any    `json:"meta,omitempty"`
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

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}
