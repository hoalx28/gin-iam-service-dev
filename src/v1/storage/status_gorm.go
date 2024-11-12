package storage

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"

	"gorm.io/gorm"
)

type gormStatusStorage struct {
	db          *gorm.DB
	userStorage UserStorage
}

func NewGormStatusStorage(appCtx config.AppContext) StatusStorage {
	return gormStatusStorage{db: appCtx.GetGormDB(), userStorage: NewGormUserStorage(appCtx)}
}

func (s gormStatusStorage) ExistByContent(content string) bool {
	err := s.db.Unscoped().Where("content = ?", content).First(&model.Status{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormStatusStorage) Save(model *model.Status) (*model.Status, exception.ServiceException) {
	owning, queriedErr := s.userStorage.FindById(model.UserID)
	if queriedErr != nil {
		return nil, exception.NewServiceException(queriedErr, constant.OwningSideNotExistedF)
	}
	model.User = owning
	savedErr := s.db.Create(model).Error
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SaveF)
	}
	return model, nil
}

func (s gormStatusStorage) FindById(id uint) (*model.Status, exception.ServiceException) {
	model := &model.Status{}
	err := s.db.Preload("User").First(model, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return model, nil
}

func (s gormStatusStorage) FindAllById(ids []uint) (*model.Statuses, exception.ServiceException) {
	models := &model.Statuses{}
	err := s.db.Preload("User").Where("id IN ?", ids).Find(models).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return models, nil
}

func (s gormStatusStorage) FindAll(page *Page) (*model.Statuses, *Paging, exception.ServiceException) {
	models := &model.Statuses{}
	paging := &Paging{}
	err := s.db.Scopes(PageScope(models, s.db, page, paging)).Preload("User").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormStatusStorage) FindAllBy(content string, page *Page) (*model.Statuses, *Paging, exception.ServiceException) {
	models := &model.Statuses{}
	paging := &Paging{}
	err := s.db.Scopes(PageContentFilterScope(models, content, s.db, page, paging)).Preload("User").Where("content LIKE ?", "%"+content+"%").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormStatusStorage) FindAllArchived(page *Page) (*model.Statuses, *Paging, exception.ServiceException) {
	models := &model.Statuses{}
	paging := &Paging{}
	err := s.db.Unscoped().Preload("User").Where("deleted_at IS NOT NULL").Scopes(PageArchivedScope(models, s.db, page, paging)).Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return models, paging, nil
}

func (s gormStatusStorage) Update(id uint, update *model.StatusUpdate) (*model.Status, exception.ServiceException) {
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

func (s gormStatusStorage) Delete(id uint) (*model.Status, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&model.Status{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
