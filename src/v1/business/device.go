package business

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type DeviceBusiness interface {
	SaveBusiness(creation *model.DeviceCreation) (*model.DeviceResponse, exception.ServiceException)
	FindByIdBusiness(id uint) (*model.DeviceResponse, exception.ServiceException)
	FindAllByIdBusiness(ids []uint) (*model.DeviceResponses, exception.ServiceException)
	FindAllBusiness(page *storage.Page) (*model.DeviceResponses, *storage.Paging, exception.ServiceException)
	FindAllByBusiness(userAgent string, page *storage.Page) (*model.DeviceResponses, *storage.Paging, exception.ServiceException)
	FindAllArchivedBusiness(page *storage.Page) (*model.DeviceResponses, *storage.Paging, exception.ServiceException)
	UpdateBusiness(id uint, update *model.DeviceUpdate) (*model.DeviceResponse, exception.ServiceException)
	DeleteBusiness(id uint) (*model.DeviceResponse, exception.ServiceException)
}
