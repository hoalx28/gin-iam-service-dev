package constant

import "net/http"

type Failed struct {
	Code       int
	Message    string
	StatusCode int
}

func newFailed(code int, message string, statusCode int) Failed {
	return Failed{Code: code, Message: message, StatusCode: statusCode}
}

var (
	UncategorizedF = newFailed(000, "uncategorized exception, service can not response.", http.StatusInternalServerError)

	EnvConfigF = newFailed(001, "can not load .env file, make sure file already existed.", http.StatusInternalServerError)
	DBConfigF  = newFailed(002, "can not established connection to database via gorm.", http.StatusInternalServerError)
	SecConfigF = newFailed(003, "can not load policy configuration.", http.StatusInternalServerError)

	RequestBodyNotReadableF   = newFailed(100, "missing or request body is not readable.", http.StatusBadRequest)
	RequestHeaderNotReadableF = newFailed(101, "missing or request header is not readable.", http.StatusBadRequest)
	RequestQueryNotReadableF  = newFailed(102, "missing or query string is not readable.", http.StatusBadRequest)
	RequestParamsNotReadableF = newFailed(103, "missing or path variable is not readable.", http.StatusBadRequest)

	FindByF                 = newFailed(200, "can not query {resource} by {criteria}.", http.StatusInternalServerError)
	FindByNoContentF        = newFailed(201, "retrieve {resource} by {criteria} return no content.", http.StatusNoContent)
	SaveF                   = newFailed(202, "can not save {resource}: try again later.", http.StatusBadRequest)
	FindByIdF               = newFailed(203, "can not retrieve {resource} by id: try again later.", http.StatusInternalServerError)
	FindByIdNoContentF      = newFailed(204, "retrieve {resource} by id return no content.", http.StatusNoContent)
	FindAllByIdF            = newFailed(205, "can not retrieve {resource} by id: try again later.", http.StatusInternalServerError)
	FindAllByIdNoContentF   = newFailed(206, "retrieve {resource} return no content.", http.StatusNoContent)
	FindAllByF              = newFailed(207, "can not retrieve {resource} by criteria: try again later.", http.StatusInternalServerError)
	FindAllByNoContentF     = newFailed(208, "retrieve {resource}s by {criteria} return no content.", http.StatusNoContent)
	FindAllArchivedF        = newFailed(209, "can not retrieve archived {resource}: try again later.", http.StatusInternalServerError)
	FindAllNoContentF       = newFailed(210, "retrieve archived {resource} return no content.", http.StatusNoContent)
	UpdateF                 = newFailed(211, "can not update {resource}: try again later.", http.StatusInternalServerError)
	DeleteF                 = newFailed(212, "can not remove {resource}: try again later.", http.StatusInternalServerError)
	AlreadyExistedF         = newFailed(213, "{resource} already existed, {identity} must be unique.", http.StatusBadRequest)
	NotExistedF             = newFailed(214, "{resource} does not existed.", http.StatusNotModified)
	OwningSideNotExistedF   = newFailed(215, "can not save {resource}: {owning side} not exists.", http.StatusBadRequest)
	OwningSideNotAvailableF = newFailed(216, "can not save {resource}: {owning side} not available.", http.StatusBadRequest)

	SignJwtTokenF     = newFailed(216, "can not sign token: ill legal claims or encrypt algorithm.", http.StatusInternalServerError)
	ParseJwtTokenF    = newFailed(216, "can not parse token: ill legal token or encrypt algorithm.", http.StatusInternalServerError)
	IllLegalJwtTokenF = newFailed(217, "ill legal token: token has been edited, expired or not publish by us.", http.StatusUnauthorized)
	JwtTokenExpiredF  = newFailed(217, "ill legal token: token has been expired.", http.StatusUnauthorized)

	SignUpF                      = newFailed(218, "can not sign up: try again later.", http.StatusInternalServerError)
	SignInF                      = newFailed(219, "can not sign in: try again later.", http.StatusInternalServerError)
	BadCredentialF               = newFailed(220, "bad credentials: username or password not match.", http.StatusUnauthorized)
	VerifiedIdentityF            = newFailed(221, "can not verify identity: try again later.", http.StatusUnauthorized)
	RetrieveProfileF             = newFailed(222, "can not retrieve profile: try again later.", http.StatusUnauthorized)
	SignOutF                     = newFailed(223, "can not sign out: token not be recalled.", http.StatusInternalServerError)
	EnsureTokenNotBadCredentialF = newFailed(224, "can not ensure token is not recall.", http.StatusUnauthorized)
	TokenBlockedF                = newFailed(225, "token has been recall: can not use this any more", http.StatusUnauthorized)
	JwtTokenNotSuitableF         = newFailed(226, "access token and refresh token are not suitable.", http.StatusUnauthorized)
	RecallJwtTokenF              = newFailed(227, "refresh token may not complete: token not be recalled.", http.StatusUnauthorized)
	RefreshTokenF                = newFailed(228, "can not refresh token: try again later.", http.StatusInternalServerError)

	UnauthorizedF = newFailed(229, "ill legal token: token has been edited, expired or not publish by us.", http.StatusUnauthorized)
	ForbiddenF    = newFailed(230, "forbidden: do not has right authority, do not f*ck with cat.", http.StatusForbidden)
)
