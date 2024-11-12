package constant

import "net/http"

type Success struct {
	Code       int32
	Message    string
	StatusCode int
}

func newSuccess(code int32, message string, statusCode int) Success {
	return Success{Code: code, Message: message, StatusCode: statusCode}
}

var (
	Save      = newSuccess(1, "{resource} has been saved.", http.StatusCreated)
	FindById  = newSuccess(1, "query {resource} by id success.", http.StatusOK)
	FindAll   = newSuccess(1, "query {resource}s success.", http.StatusOK)
	FindAllBy = newSuccess(1, "query {resource} by {criteria} success.", http.StatusOK)
	Update    = newSuccess(1, "{resource} has been updated.", http.StatusOK)
	Delete    = newSuccess(1, "{resource} has been deleted.", http.StatusOK)
)
