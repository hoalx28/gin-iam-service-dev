package storage

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
)

type DeviceStorage interface {
	ExistByIpAddress(ipAddress string) bool
	Save(model *model.Device) (*model.Device, exception.ServiceException)
	FindById(id uint) (*model.Device, exception.ServiceException)
	FindAllById(ids []uint) (*model.Devices, exception.ServiceException)
	FindAll(page *Page) (*model.Devices, *Paging, exception.ServiceException)
	FindAllBy(userAgent string, page *Page) (*model.Devices, *Paging, exception.ServiceException)
	FindAllArchived(page *Page) (*model.Devices, *Paging, exception.ServiceException)
	Update(id uint, update *model.DeviceUpdate) (*model.Device, exception.ServiceException)
	Delete(id uint) (*model.Device, exception.ServiceException)
}
