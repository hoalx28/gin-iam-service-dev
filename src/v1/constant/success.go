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

	SignUp          = newSuccess(1, "sign up success, enjoy.", http.StatusCreated)
	SignIn          = newSuccess(1, "sign in success, enjoy.", http.StatusOK)
	VerifyIdentity  = newSuccess(1, "identity has been verified, enjoy.", http.StatusOK)
	RetrieveProfile = newSuccess(1, "retrieve profile success, enjoy.", http.StatusOK)
	SignOut         = newSuccess(1, "sign out success, enjoy.", http.StatusOK)
	RefreshToken    = newSuccess(1, "refresh token success, enjoy.", http.StatusOK)
)
