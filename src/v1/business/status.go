package business

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type StatusBusiness interface {
	SaveBusiness(creation *model.StatusCreation) (*model.StatusResponse, exception.ServiceException)
	FindByIdBusiness(id uint) (*model.StatusResponse, exception.ServiceException)
	FindAllByIdBusiness(ids []uint) (*model.StatusResponses, exception.ServiceException)
	FindAllBusiness(page *storage.Page) (*model.StatusResponses, *storage.Paging, exception.ServiceException)
	FindAllByBusiness(content string, page *storage.Page) (*model.StatusResponses, *storage.Paging, exception.ServiceException)
	FindAllArchivedBusiness(page *storage.Page) (*model.StatusResponses, *storage.Paging, exception.ServiceException)
	UpdateBusiness(id uint, update *model.StatusUpdate) (*model.StatusResponse, exception.ServiceException)
	DeleteBusiness(id uint) (*model.StatusResponse, exception.ServiceException)
}
