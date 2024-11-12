package storage

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"

	"gorm.io/gorm"
)

type gormDeviceStorage struct {
	db          *gorm.DB
	userStorage UserStorage
}

func NewGormDeviceStorage(appCtx config.AppContext) DeviceStorage {
	return gormDeviceStorage{db: appCtx.GetGormDB(), userStorage: NewGormUserStorage(appCtx)}
}

func (s gormDeviceStorage) ExistByIpAddress(ipAddress string) bool {
	err := s.db.Unscoped().Where("ip_address = ?", ipAddress).First(&model.Device{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormDeviceStorage) Save(model *model.Device) (*model.Device, exception.ServiceException) {
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

func (s gormDeviceStorage) FindById(id uint) (*model.Device, exception.ServiceException) {
	model := &model.Device{}
	err := s.db.First(model, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return model, nil
}

func (s gormDeviceStorage) FindAllById(ids []uint) (*model.Devices, exception.ServiceException) {
	models := &model.Devices{}
	err := s.db.Where("id IN ?", ids).Find(models).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return models, nil
}

func (s gormDeviceStorage) FindAll(page *Page) (*model.Devices, *Paging, exception.ServiceException) {
	models := &model.Devices{}
	paging := &Paging{}
	err := s.db.Scopes(PageScope(models, s.db, page, paging)).Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormDeviceStorage) FindAllBy(userAgent string, page *Page) (*model.Devices, *Paging, exception.ServiceException) {
	models := &model.Devices{}
	paging := &Paging{}
	err := s.db.Scopes(PageUserAgentFilterScope(models, userAgent, s.db, page, paging)).Where("user_agent LIKE ?", "%"+userAgent+"%").Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return models, paging, nil
}

func (s gormDeviceStorage) FindAllArchived(page *Page) (*model.Devices, *Paging, exception.ServiceException) {
	models := &model.Devices{}
	paging := &Paging{}
	err := s.db.Unscoped().Where("deleted_at IS NOT NULL").Scopes(PageArchivedScope(models, s.db, page, paging)).Find(models).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return models, paging, nil
}

func (s gormDeviceStorage) Update(id uint, update *model.DeviceUpdate) (*model.Device, exception.ServiceException) {
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

func (s gormDeviceStorage) Delete(id uint) (*model.Device, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&model.Device{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
