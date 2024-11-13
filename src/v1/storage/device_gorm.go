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

type gormDeviceS struct {
	db    *gorm.DB
	UserS abstraction.UserS
}

func NewGormDeviceS(appCtx config.AppContext) abstraction.DeviceS {
	return gormDeviceS{db: appCtx.GetGormDB(), UserS: NewGormUserS(appCtx)}
}

func (s gormDeviceS) ExistByIpAddress(ipAddress string) bool {
	err := s.db.Unscoped().Where("ip_address = ?", ipAddress).First(&domain.Device{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormDeviceS) Save(domain *domain.Device) (*domain.Device, exception.ServiceException) {
	owning, queriedErr := s.UserS.FindById(domain.UserID)
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

func (s gormDeviceS) FindById(id uint) (*domain.Device, exception.ServiceException) {
	domain := &domain.Device{}
	err := s.db.First(domain, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return domain, nil
}

func (s gormDeviceS) FindAllById(ids []uint) (*domain.Devices, exception.ServiceException) {
	domains := &domain.Devices{}
	err := s.db.Where("id IN ?", ids).Find(domains).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return domains, nil
}

func (s gormDeviceS) FindAll(page dto.Page) (*domain.Devices, *dto.Paging, exception.ServiceException) {
	domains := &domain.Devices{}
	paging := &dto.Paging{}
	err := s.db.Scopes(PageScope(domains, s.db, page, paging)).Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormDeviceS) FindAllBy(userAgent string, page dto.Page) (*domain.Devices, *dto.Paging, exception.ServiceException) {
	domains := &domain.Devices{}
	paging := &dto.Paging{}
	err := s.db.Scopes(ConditionPageScope(domains, map[string]string{"UserAgent": userAgent}, s.db, page, paging)).Where("user_agent LIKE ?", "%"+userAgent+"%").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormDeviceS) FindAllArchived(page dto.Page) (*domain.Devices, *dto.Paging, exception.ServiceException) {
	domains := &domain.Devices{}
	paging := &dto.Paging{}
	err := s.db.Unscoped().Where("deleted_at IS NOT NULL").Scopes(ArchivedPageScope(domains, s.db, page, paging)).Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return domains, paging, nil
}

func (s gormDeviceS) Update(id uint, update *domain.DeviceUpdate) (*domain.Device, exception.ServiceException) {
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

func (s gormDeviceS) Delete(id uint) (*domain.Device, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&domain.Device{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
