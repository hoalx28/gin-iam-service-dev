package response

type response[T any] struct {
	Timestamp  int64  `json:"timestamp,omitempty"`
	Code       int32  `json:"code,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Payload    T      `json:"payload,omitempty"`
}

func NewResponse[T any](timestamp int64, code int32, statusCode int, message string, payload T) response[T] {
	return response[T]{Timestamp: timestamp, Code: code, StatusCode: statusCode, Message: message, Payload: payload}
}
