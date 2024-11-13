package business

import (
	"iam/src/v1/exception"
	"iam/src/v1/model/dto"
	"iam/src/v1/token"
)

type AuthBusiness interface {
	SignUp(request dto.RegisterRequest) (*dto.CredentialResponse, exception.ServiceException)
	SignIn(request dto.CredentialRequest) (*dto.CredentialResponse, exception.ServiceException)
	Identity(accessToken string) (*token.AuthClaims, exception.ServiceException)
	Me(accessToken string) (*dto.RegisterResponse, exception.ServiceException)
	SignOut(accessToken string) (*uint, exception.ServiceException)
	Refresh(accessToken string, refreshToken string) (*dto.CredentialResponse, exception.ServiceException)
}
