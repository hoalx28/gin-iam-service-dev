package business

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type UserBusiness interface {
	SaveBusiness(creation *model.UserCreation) (*model.UserResponse, exception.ServiceException)
	FindByIdBusiness(id uint) (*model.UserResponse, exception.ServiceException)
	FindAllByIdBusiness(ids []uint) (*model.UserResponses, exception.ServiceException)
	FindAllBusiness(page *storage.Page) (*model.UserResponses, *storage.Paging, exception.ServiceException)
	FindAllByBusiness(name string, page *storage.Page) (*model.UserResponses, *storage.Paging, exception.ServiceException)
	FindAllArchivedBusiness(page *storage.Page) (*model.UserResponses, *storage.Paging, exception.ServiceException)
	UpdateBusiness(id uint, update *model.UserUpdate) (*model.UserResponse, exception.ServiceException)
	DeleteBusiness(id uint) (*model.UserResponse, exception.ServiceException)
}
