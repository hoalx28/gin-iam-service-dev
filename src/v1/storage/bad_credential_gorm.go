package storage

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/exception"

	"gorm.io/gorm"
)

type gormBadCredentialS struct {
	db *gorm.DB
}

func NewGormBadCredentialS(appCtx config.AppContext) abstraction.BadCredentialS {
	return gormBadCredentialS{db: appCtx.GetGormDB()}
}

func (s gormBadCredentialS) Save(domain *domain.BadCredential) (*domain.BadCredential, exception.ServiceException) {
	err := s.db.Create(domain).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.SaveF)
	}
	return domain, nil
}

func (s gormBadCredentialS) FindByAccessTokenId(accessToken string) (*domain.BadCredential, exception.ServiceException) {
	domain := &domain.BadCredential{}
	err := s.db.Unscoped().Where("access_token_id = ?", accessToken).First(domain).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return domain, nil
}
