package storage

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"

	"gorm.io/gorm"
)

type gormRoleStorage struct {
	db               *gorm.DB
	privilegeStorage PrivilegeStorage
}

func NewGormRoleStorage(appCtx config.AppContext) RoleStorage {
	return gormRoleStorage{db: appCtx.GetGormDB(), privilegeStorage: NewGormPrivilegeStorage(appCtx)}
}

func (s gormRoleStorage) ExistByName(name string) bool {
	err := s.db.Unscoped().Where("name = ?", name).First(&model.Role{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormRoleStorage) Save(model *model.Role) (*model.Role, exception.ServiceException) {
	owning, queriedErr := s.privilegeStorage.FindAllById(model.PrivilegeIds)
	isOwningSideNotExisted := queriedErr != nil || len(*owning) == 0
	if isOwningSideNotExisted {
		return nil, exception.NewServiceException(queriedErr, constant.OwningSideNotExistedF)
	}
	model.Privileges = *owning
	savedErr := s.db.Create(model).Error
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SaveF)
	}
	return model, nil
}

func (s gormRoleStorage) FindById(id uint) (*model.Role, exception.ServiceException) {
	model := &model.Role{}
	err := s.db.Preload("Privileges").First(model, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return model, nil
}

func (s gormRoleStorage) FindAllById(ids []uint) (*model.Roles, exception.ServiceException) {
	models := &model.Roles{}
	err := s.db.Preload("Privileges").Where("id IN ?", ids).Find(models).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return models, nil
}

func (s gormRoleStorage) FindAll(page *Page) (*model.Roles, *Paging, exception.ServiceException) {
	models := &model.Roles{}
	paging := &Paging{}
	err := s.db.Scopes(PageScope(models, s.db, page, paging)).Preload("Privileges").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormRoleStorage) FindAllBy(name string, page *Page) (*model.Roles, *Paging, exception.ServiceException) {
	models := &model.Roles{}
	paging := &Paging{}
	err := s.db.Scopes(PageNameFilterScope(models, name, s.db, page, paging)).Preload("Privileges").Where("name LIKE ?", "%"+name+"%").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormRoleStorage) FindAllArchived(page *Page) (*model.Roles, *Paging, exception.ServiceException) {
	models := &model.Roles{}
	paging := &Paging{}
	err := s.db.Unscoped().Preload("Privileges").Where("deleted_at IS NOT NULL").Scopes(PageArchivedScope(models, s.db, page, paging)).Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return models, paging, nil
}

func (s gormRoleStorage) Update(id uint, update *model.RoleUpdate) (*model.Role, exception.ServiceException) {
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

func (s gormRoleStorage) Delete(id uint) (*model.Role, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&model.Role{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
