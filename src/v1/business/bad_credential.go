package business

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
)

type BadCredentialBusiness interface {
	SaveBusiness(creation *model.BadCredentialCreation) (*model.BadCredentialResponse, exception.ServiceException)
	FindByAccessTokenIdBusiness(accessTokenId string) (*model.BadCredentialResponse, exception.ServiceException)
}
