package storage

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"

	"gorm.io/gorm"
)

type gormUserStorage struct {
	db          *gorm.DB
	roleStorage RoleStorage
}

func NewGormUserStorage(appCtx config.AppContext) UserStorage {
	return gormUserStorage{db: appCtx.GetGormDB(), roleStorage: NewGormRoleStorage(appCtx)}
}

func (s gormUserStorage) ExistByUsername(username string) bool {
	err := s.db.Unscoped().Where("username = ?", username).First(&model.User{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormUserStorage) Save(model *model.User) (*model.User, exception.ServiceException) {
	owning, queriedErr := s.roleStorage.FindAllById(model.RoleIds)
	isOwningSideNotExisted := queriedErr != nil || len(*owning) == 0
	if isOwningSideNotExisted {
		return nil, exception.NewServiceException(queriedErr, constant.OwningSideNotExistedF)
	}
	model.Roles = *owning
	savedErr := s.db.Create(model).Error
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SaveF)
	}
	return model, nil
}

func (s gormUserStorage) FindById(id uint) (*model.User, exception.ServiceException) {
	model := &model.User{}
	err := s.db.Preload("Roles").First(model, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return model, nil
}

func (s gormUserStorage) FindAllById(ids []uint) (*model.Users, exception.ServiceException) {
	models := &model.Users{}
	err := s.db.Preload("Roles").Where("id IN ?", ids).Find(models).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return models, nil
}

func (s gormUserStorage) FindAll(page *Page) (*model.Users, *Paging, exception.ServiceException) {
	models := &model.Users{}
	paging := &Paging{}
	err := s.db.Scopes(PageScope(models, s.db, page, paging)).Preload("Roles").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormUserStorage) FindAllBy(username string, page *Page) (*model.Users, *Paging, exception.ServiceException) {
	models := &model.Users{}
	paging := &Paging{}
	err := s.db.Scopes(PageUsernameFilterScope(models, username, s.db, page, paging)).Preload("Roles").Where("username LIKE ?", "%"+username+"%").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormUserStorage) FindAllArchived(page *Page) (*model.Users, *Paging, exception.ServiceException) {
	models := &model.Users{}
	paging := &Paging{}
	err := s.db.Unscoped().Preload("Roles").Where("deleted_at IS NOT NULL").Scopes(PageArchivedScope(models, s.db, page, paging)).Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return models, paging, nil
}

func (s gormUserStorage) Update(id uint, update *model.UserUpdate) (*model.User, exception.ServiceException) {
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

func (s gormUserStorage) Delete(id uint) (*model.User, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
