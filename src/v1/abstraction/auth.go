package abstraction

import (
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
)

type AuthB interface {
	SignUp(request dto.RegisterRequest) (*dto.CredentialResponse, exception.ServiceException)
	SignIn(request dto.CredentialRequest) (*dto.CredentialResponse, exception.ServiceException)
	Identity(accessToken string) (*dto.AuthClaims, exception.ServiceException)
	Me(accessToken string) (*dto.RegisterResponse, exception.ServiceException)
	SignOut(accessToken string) (*uint, exception.ServiceException)
	Refresh(accessToken string, refreshToken string) (*dto.CredentialResponse, exception.ServiceException)
}
