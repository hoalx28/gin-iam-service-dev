package token

import (
	"iam/src/v1/exception"
)

type TokenAuthProvider[T any, U any] interface {
	BuildScope(user U) string
	Sign(user U, expiredTime int, secretKey string, id string, referId string) (*string, exception.ServiceException)
	Verify(token string, secretKey string) (*T, exception.ServiceException)
	EnsureNotBadCredential(token string) (*AuthClaims, exception.ServiceException)
}
