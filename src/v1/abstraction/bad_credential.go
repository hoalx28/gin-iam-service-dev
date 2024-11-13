package abstraction

import (
	"iam/src/v1/domain"
	"iam/src/v1/exception"
)

type BadCredentialS interface {
	Save(domain *domain.BadCredential) (*domain.BadCredential, exception.ServiceException)
	FindByAccessTokenId(accessToken string) (*domain.BadCredential, exception.ServiceException)
}

type BadCredentialB interface {
	SaveB(creation *domain.BadCredentialCreation) (*domain.BadCredentialResponse, exception.ServiceException)
	FindByAccessTokenIdB(accessTokenId string) (*domain.BadCredentialResponse, exception.ServiceException)
}
