package abstraction

import (
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
)

type RoleS interface {
	ExistByName(name string) bool
	Save(domain *domain.Role) (*domain.Role, exception.ServiceException)
	FindById(id uint) (*domain.Role, exception.ServiceException)
	FindAllById(ids []uint) (*domain.Roles, exception.ServiceException)
	FindAll(page dto.Page) (*domain.Roles, *dto.Paging, exception.ServiceException)
	FindAllBy(name string, page dto.Page) (*domain.Roles, *dto.Paging, exception.ServiceException)
	FindAllArchived(page dto.Page) (*domain.Roles, *dto.Paging, exception.ServiceException)
	Update(id uint, update *domain.RoleUpdate) (*domain.Role, exception.ServiceException)
	Delete(id uint) (*domain.Role, exception.ServiceException)
}

type RoleB interface {
	SaveB(creation *domain.RoleCreation) (*domain.RoleResponse, exception.ServiceException)
	FindByIdB(id uint) (*domain.RoleResponse, exception.ServiceException)
	FindAllByIdB(ids []uint) (*domain.RoleResponses, exception.ServiceException)
	FindAllB(page dto.Page) (*domain.RoleResponses, *dto.Paging, exception.ServiceException)
	FindAllByB(name string, page dto.Page) (*domain.RoleResponses, *dto.Paging, exception.ServiceException)
	FindAllArchivedB(page dto.Page) (*domain.RoleResponses, *dto.Paging, exception.ServiceException)
	UpdateB(id uint, update *domain.RoleUpdate) (*domain.RoleResponse, exception.ServiceException)
	DeleteB(id uint) (*domain.RoleResponse, exception.ServiceException)
}
