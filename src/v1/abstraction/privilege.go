package abstraction

import (
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
)

type PrivilegeS interface {
	ExistByName(name string) bool
	Save(domain *domain.Privilege) (*domain.Privilege, exception.ServiceException)
	FindById(id uint) (*domain.Privilege, exception.ServiceException)
	FindAllById(ids []uint) (*domain.Privileges, exception.ServiceException)
	FindAll(page dto.Page) (*domain.Privileges, *dto.Paging, exception.ServiceException)
	FindAllBy(name string, page dto.Page) (*domain.Privileges, *dto.Paging, exception.ServiceException)
	FindAllArchived(page dto.Page) (*domain.Privileges, *dto.Paging, exception.ServiceException)
	Update(id uint, update *domain.PrivilegeUpdate) (*domain.Privilege, exception.ServiceException)
	Delete(id uint) (*domain.Privilege, exception.ServiceException)
}

type PrivilegeB interface {
	SaveB(creation *domain.PrivilegeCreation) (*domain.PrivilegeResponse, exception.ServiceException)
	FindByIdB(id uint) (*domain.PrivilegeResponse, exception.ServiceException)
	FindAllByIdB(ids []uint) (*domain.PrivilegeResponses, exception.ServiceException)
	FindAllB(page dto.Page) (*domain.PrivilegeResponses, *dto.Paging, exception.ServiceException)
	FindAllByB(name string, page dto.Page) (*domain.PrivilegeResponses, *dto.Paging, exception.ServiceException)
	FindAllArchivedB(page dto.Page) (*domain.PrivilegeResponses, *dto.Paging, exception.ServiceException)
	UpdateB(id uint, update *domain.PrivilegeUpdate) (*domain.PrivilegeResponse, exception.ServiceException)
	DeleteB(id uint) (*domain.PrivilegeResponse, exception.ServiceException)
}
