package abstraction

import (
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
)

type StatusS interface {
	ExistByContent(content string) bool
	Save(domain *domain.Status) (*domain.Status, exception.ServiceException)
	FindById(id uint) (*domain.Status, exception.ServiceException)
	FindAllById(ids []uint) (*domain.Statuses, exception.ServiceException)
	FindAll(page dto.Page) (*domain.Statuses, *dto.Paging, exception.ServiceException)
	FindAllBy(content string, page dto.Page) (*domain.Statuses, *dto.Paging, exception.ServiceException)
	FindAllArchived(page dto.Page) (*domain.Statuses, *dto.Paging, exception.ServiceException)
	Update(id uint, update *domain.StatusUpdate) (*domain.Status, exception.ServiceException)
	Delete(id uint) (*domain.Status, exception.ServiceException)
}

type StatusB interface {
	SaveB(creation *domain.StatusCreation) (*domain.StatusResponse, exception.ServiceException)
	FindByIdB(id uint) (*domain.StatusResponse, exception.ServiceException)
	FindAllByIdB(ids []uint) (*domain.StatusResponses, exception.ServiceException)
	FindAllB(page dto.Page) (*domain.StatusResponses, *dto.Paging, exception.ServiceException)
	FindAllByB(content string, page dto.Page) (*domain.StatusResponses, *dto.Paging, exception.ServiceException)
	FindAllArchivedB(page dto.Page) (*domain.StatusResponses, *dto.Paging, exception.ServiceException)
	UpdateB(id uint, update *domain.StatusUpdate) (*domain.StatusResponse, exception.ServiceException)
	DeleteB(id uint) (*domain.StatusResponse, exception.ServiceException)
}
