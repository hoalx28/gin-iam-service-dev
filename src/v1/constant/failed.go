package constant

import "net/http"

type Failed struct {
	Code       int32
	Message    string
	StatusCode int
}

func newFailed(code int32, message string, statusCode int) Failed {
	return Failed{Code: code, Message: message, StatusCode: statusCode}
}

var (
	EnvConfigF = newFailed(001, "can not load .env file, make sure file already existed.", http.StatusInternalServerError)
	DBConfigF  = newFailed(002, "can not established connection to database via gorm.", http.StatusInternalServerError)

	RequestBodyNotReadableF   = newFailed(102, "missing or request body is not readable.", http.StatusBadRequest)
	RequestQueryNotReadableF  = newFailed(102, "missing or query string is not readable.", http.StatusBadRequest)
	RequestParamsNotReadableF = newFailed(102, "missing or path variable is not readable.", http.StatusBadRequest)

	FindByF                 = newFailed(200, "can not query {resource} by {criteria}.", http.StatusInternalServerError)
	SaveF                   = newFailed(201, "can not save {resource}: try again later.", http.StatusBadRequest)
	FindByIdF               = newFailed(202, "can not retrieve {resource} by id: try again later.", http.StatusInternalServerError)
	FindByIdNoContentF      = newFailed(203, "retrieve {resource} by id return no content.", http.StatusNoContent)
	FindAllByIdF            = newFailed(204, "can not retrieve {resource} by id: try again later.", http.StatusInternalServerError)
	FindAllByIdNoContentF   = newFailed(205, "retrieve {resource} return no content.", http.StatusNoContent)
	FindAllByF              = newFailed(206, "can not retrieve {resource} by criteria: try again later.", http.StatusInternalServerError)
	FindAllByNoContentF     = newFailed(207, "retrieve {resource}s by {criteria} return no content.", http.StatusNoContent)
	FindAllArchivedF        = newFailed(208, "can not retrieve archived {resource}: try again later.", http.StatusInternalServerError)
	FindAllNoContentF       = newFailed(209, "retrieve archived {resource} return no content.", http.StatusNoContent)
	UpdateF                 = newFailed(210, "can not update {resource}: try again later.", http.StatusInternalServerError)
	DeleteF                 = newFailed(211, "can not remove {resource}: try again later.", http.StatusInternalServerError)
	AlreadyExistedF         = newFailed(212, "{resource} already existed, {identity} must be unique.", http.StatusBadRequest)
	NotExistedF             = newFailed(213, "{resource} does not existed.", http.StatusNotModified)
	OwningSideNotExistedF   = newFailed(214, "can not save {resource}: {owning side} not exists.", http.StatusBadRequest)
	OwningSideNotAvailableF = newFailed(214, "can not save {resource}: {owning side} not available.", http.StatusBadRequest)
)
