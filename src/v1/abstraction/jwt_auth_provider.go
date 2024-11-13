package abstraction

import (
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
)

type JwtAuthProvider[T any, U any] interface {
	BuildScope(user U) string
	Sign(user U, expiredTime int, secretKey string, id string, referId string) (*string, exception.ServiceException)
	Verify(token string, secretKey string) (*T, exception.ServiceException)
	EnsureNotBadCredential(token string) (*T, exception.ServiceException)
	GetConstant() dto.Token
}
