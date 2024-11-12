package business

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type PrivilegeBusiness interface {
	SaveBusiness(creation *model.PrivilegeCreation) (*model.PrivilegeResponse, exception.ServiceException)
	FindByIdBusiness(id uint) (*model.PrivilegeResponse, exception.ServiceException)
	FindAllByIdBusiness(ids []uint) (*model.PrivilegeResponses, exception.ServiceException)
	FindAllBusiness(page *storage.Page) (*model.PrivilegeResponses, *storage.Paging, exception.ServiceException)
	FindAllByBusiness(name string, page *storage.Page) (*model.PrivilegeResponses, *storage.Paging, exception.ServiceException)
	FindAllArchivedBusiness(page *storage.Page) (*model.PrivilegeResponses, *storage.Paging, exception.ServiceException)
	UpdateBusiness(id uint, update *model.PrivilegeUpdate) (*model.PrivilegeResponse, exception.ServiceException)
	DeleteBusiness(id uint) (*model.PrivilegeResponse, exception.ServiceException)
}
