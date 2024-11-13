package storage

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"

	"gorm.io/gorm"
)

type gormStatusS struct {
	db    *gorm.DB
	userS abstraction.UserS
}

func NewGormStatusS(appCtx config.AppContext) abstraction.StatusS {
	return gormStatusS{db: appCtx.GetGormDB(), userS: NewGormUserS(appCtx)}
}

func (s gormStatusS) ExistByContent(content string) bool {
	err := s.db.Unscoped().Where("content = ?", content).First(&domain.Status{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormStatusS) Save(domain *domain.Status) (*domain.Status, exception.ServiceException) {
	owning, queriedErr := s.userS.FindById(domain.UserID)
	if queriedErr != nil {
		return nil, exception.NewServiceException(queriedErr, constant.OwningSideNotExistedF)
	}
	domain.User = owning
	savedErr := s.db.Create(domain).Error
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SaveF)
	}
	return domain, nil
}

func (s gormStatusS) FindById(id uint) (*domain.Status, exception.ServiceException) {
	domain := &domain.Status{}
	err := s.db.Preload("User").First(domain, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return domain, nil
}

func (s gormStatusS) FindAllById(ids []uint) (*domain.Statuses, exception.ServiceException) {
	domains := &domain.Statuses{}
	err := s.db.Preload("User").Where("id IN ?", ids).Find(domains).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return domains, nil
}

func (s gormStatusS) FindAll(page dto.Page) (*domain.Statuses, *dto.Paging, exception.ServiceException) {
	domains := &domain.Statuses{}
	paging := &dto.Paging{}
	err := s.db.Scopes(PageScope(domains, s.db, page, paging)).Preload("User").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormStatusS) FindAllBy(content string, page dto.Page) (*domain.Statuses, *dto.Paging, exception.ServiceException) {
	domains := &domain.Statuses{}
	paging := &dto.Paging{}
	err := s.db.Scopes(ConditionPageScope(domains, map[string]string{"Content": content}, s.db, page, paging)).Preload("User").Where("content LIKE ?", "%"+content+"%").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormStatusS) FindAllArchived(page dto.Page) (*domain.Statuses, *dto.Paging, exception.ServiceException) {
	domains := &domain.Statuses{}
	paging := &dto.Paging{}
	err := s.db.Unscoped().Preload("User").Where("deleted_at IS NOT NULL").Scopes(ArchivedPageScope(domains, s.db, page, paging)).Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return domains, paging, nil
}

func (s gormStatusS) Update(id uint, update *domain.StatusUpdate) (*domain.Status, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Updates(update).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.UpdateF)
	}
	return old, nil
}

func (s gormStatusS) Delete(id uint) (*domain.Status, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&domain.Status{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
