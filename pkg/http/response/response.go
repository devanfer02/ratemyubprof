package response

type Response struct {
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

