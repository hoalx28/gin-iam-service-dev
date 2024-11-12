package storage

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"

	"gorm.io/gorm"
)

type gormPrivilegeStorage struct {
	db *gorm.DB
}

func NewGormPrivilegeStorage(appCtx config.AppContext) PrivilegeStorage {
	return gormPrivilegeStorage{db: appCtx.GetGormDB()}
}

func (s gormPrivilegeStorage) ExistByName(name string) bool {
	err := s.db.Unscoped().Where("name = ?", name).First(&model.Privilege{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormPrivilegeStorage) Save(model *model.Privilege) (*model.Privilege, exception.ServiceException) {
	err := s.db.Create(model).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.SaveF)
	}
	return model, nil
}

func (s gormPrivilegeStorage) FindById(id uint) (*model.Privilege, exception.ServiceException) {
	model := &model.Privilege{}
	err := s.db.First(model, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return model, nil
}

func (s gormPrivilegeStorage) FindAllById(ids []uint) (*model.Privileges, exception.ServiceException) {
	models := &model.Privileges{}
	err := s.db.Where("id IN ?", ids).Find(models).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return models, nil
}

func (s gormPrivilegeStorage) FindAll(page *Page) (*model.Privileges, *Paging, exception.ServiceException) {
	models := &model.Privileges{}
	paging := &Paging{}
	err := s.db.Scopes(PageScope(models, s.db, page, paging)).Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormPrivilegeStorage) FindAllBy(name string, page *Page) (*model.Privileges, *Paging, exception.ServiceException) {
	models := &model.Privileges{}
	paging := &Paging{}
	err := s.db.Scopes(PageNameFilterScope(models, name, s.db, page, paging)).Where("name LIKE ?", "%"+name+"%").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormPrivilegeStorage) FindAllArchived(page *Page) (*model.Privileges, *Paging, exception.ServiceException) {
	models := &model.Privileges{}
	paging := &Paging{}
	err := s.db.Unscoped().Where("deleted_at IS NOT NULL").Scopes(PageArchivedScope(models, s.db, page, paging)).Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return models, paging, nil
}

func (s gormPrivilegeStorage) Update(id uint, update *model.PrivilegeUpdate) (*model.Privilege, exception.ServiceException) {
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

func (s gormPrivilegeStorage) Delete(id uint) (*model.Privilege, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&model.Privilege{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
