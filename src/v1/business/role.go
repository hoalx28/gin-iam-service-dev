package business

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type RoleBusiness interface {
	SaveBusiness(creation *model.RoleCreation) (*model.RoleResponse, exception.ServiceException)
	FindByIdBusiness(id uint) (*model.RoleResponse, exception.ServiceException)
	FindAllByIdBusiness(ids []uint) (*model.RoleResponses, exception.ServiceException)
	FindAllBusiness(page *storage.Page) (*model.RoleResponses, *storage.Paging, exception.ServiceException)
	FindAllByBusiness(name string, page *storage.Page) (*model.RoleResponses, *storage.Paging, exception.ServiceException)
	FindAllArchivedBusiness(page *storage.Page) (*model.RoleResponses, *storage.Paging, exception.ServiceException)
	UpdateBusiness(id uint, update *model.RoleUpdate) (*model.RoleResponse, exception.ServiceException)
	DeleteBusiness(id uint) (*model.RoleResponse, exception.ServiceException)
}
