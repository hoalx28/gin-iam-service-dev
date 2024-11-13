package business

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/domain"
	"iam/src/v1/exception"
	"iam/src/v1/storage"
)

type gormBadCredentialB struct {
	storage abstraction.BadCredentialS
}

func NewGormBadCredentialB(appCtx config.AppContext) abstraction.BadCredentialB {
	storage := storage.NewGormBadCredentialS(appCtx)
	return gormBadCredentialB{storage: storage}
}

func (b gormBadCredentialB) SaveB(creation *domain.BadCredentialCreation) (*domain.BadCredentialResponse, exception.ServiceException) {
	model := creation.AsModel()
	saved, saveErr := b.storage.Save(&model)
	if saveErr != nil {
		return nil, saveErr
	}
	response := saved.AsResponse()
	return &response, nil
}

func (b gormBadCredentialB) FindByAccessTokenIdB(accessTokenId string) (*domain.BadCredentialResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindByAccessTokenId(accessTokenId)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return &response, nil
}
