package storage

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
)

type BadCredentialStorage interface {
	Save(model *model.BadCredential) (*model.BadCredential, exception.ServiceException)
	FindByAccessTokenId(accessToken string) (*model.BadCredential, exception.ServiceException)
}
