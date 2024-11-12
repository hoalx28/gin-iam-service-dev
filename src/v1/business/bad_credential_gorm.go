package business

import (
	"iam/src/v1/config"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type gormBadCredentialBusiness struct {
	storage storage.BadCredentialStorage
}

func NewGormBadCredentialBusiness(appCtx config.AppContext) BadCredentialBusiness {
	storage := storage.NewGormBadCredentialStorage(appCtx)
	return gormBadCredentialBusiness{storage: storage}
}

func (b gormBadCredentialBusiness) SaveBusiness(creation *model.BadCredentialCreation) (*model.BadCredentialResponse, exception.ServiceException) {
	model := creation.AsModel()
	saved, saveErr := b.storage.Save(model)
	if saveErr != nil {
		return nil, saveErr
	}
	response := saved.AsResponse()
	return response, nil
}

func (b gormBadCredentialBusiness) FindByAccessTokenIdBusiness(accessTokenId string) (*model.BadCredentialResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindByAccessTokenId(accessTokenId)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return response, nil
}
