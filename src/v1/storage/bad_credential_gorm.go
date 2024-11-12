package storage

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"

	"gorm.io/gorm"
)

type gormBadCredentialStorage struct {
	db *gorm.DB
}

func NewGormBadCredentialStorage(appCtx config.AppContext) BadCredentialStorage {
	return gormBadCredentialStorage{db: appCtx.GetGormDB()}
}

func (s gormBadCredentialStorage) Save(model *model.BadCredential) (*model.BadCredential, exception.ServiceException) {
	err := s.db.Create(model).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.SaveF)
	}
	return model, nil
}

func (s gormBadCredentialStorage) FindByAccessTokenId(accessToken string) (*model.BadCredential, exception.ServiceException) {
	model := &model.BadCredential{}
	err := s.db.Unscoped().Where("access_token_id = ?", accessToken).First(model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return model, nil
}
