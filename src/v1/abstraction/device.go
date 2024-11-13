package abstraction

import (
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
)

type DeviceS interface {
	ExistByIpAddress(ipAddress string) bool
	Save(domain *domain.Device) (*domain.Device, exception.ServiceException)
	FindById(id uint) (*domain.Device, exception.ServiceException)
	FindAllById(ids []uint) (*domain.Devices, exception.ServiceException)
	FindAll(page dto.Page) (*domain.Devices, *dto.Paging, exception.ServiceException)
	FindAllBy(userAgent string, page dto.Page) (*domain.Devices, *dto.Paging, exception.ServiceException)
	FindAllArchived(page dto.Page) (*domain.Devices, *dto.Paging, exception.ServiceException)
	Update(id uint, update *domain.DeviceUpdate) (*domain.Device, exception.ServiceException)
	Delete(id uint) (*domain.Device, exception.ServiceException)
}

type DeviceB interface {
	SaveB(creation *domain.DeviceCreation) (*domain.DeviceResponse, exception.ServiceException)
	FindByIdB(id uint) (*domain.DeviceResponse, exception.ServiceException)
	FindAllByIdB(ids []uint) (*domain.DeviceResponses, exception.ServiceException)
	FindAllB(page dto.Page) (*domain.DeviceResponses, *dto.Paging, exception.ServiceException)
	FindAllByB(userAgent string, page dto.Page) (*domain.DeviceResponses, *dto.Paging, exception.ServiceException)
	FindAllArchivedB(page dto.Page) (*domain.DeviceResponses, *dto.Paging, exception.ServiceException)
	UpdateB(id uint, update *domain.DeviceUpdate) (*domain.DeviceResponse, exception.ServiceException)
	DeleteB(id uint) (*domain.DeviceResponse, exception.ServiceException)
}
