package abstraction

import (
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
)

type UserS interface {
	ExistByUsername(username string) bool
	Save(domain *domain.User) (*domain.User, exception.ServiceException)
	FindById(id uint) (*domain.User, exception.ServiceException)
	FindByUsername(username string) (*domain.User, exception.ServiceException)
	FindAllById(ids []uint) (*domain.Users, exception.ServiceException)
	FindAll(page dto.Page) (*domain.Users, *dto.Paging, exception.ServiceException)
	FindAllBy(username string, page dto.Page) (*domain.Users, *dto.Paging, exception.ServiceException)
	FindAllArchived(page dto.Page) (*domain.Users, *dto.Paging, exception.ServiceException)
	Update(id uint, update *domain.UserUpdate) (*domain.User, exception.ServiceException)
	Delete(id uint) (*domain.User, exception.ServiceException)
}

type UserB interface {
	SaveB(creation *domain.UserCreation) (*domain.UserResponse, exception.ServiceException)
	FindByIdB(id uint) (*domain.UserResponse, exception.ServiceException)
	FindByUsernameB(username string) (*domain.UserResponse, exception.ServiceException)
	FindAllByIdB(ids []uint) (*domain.UserResponses, exception.ServiceException)
	FindAllB(page dto.Page) (*domain.UserResponses, *dto.Paging, exception.ServiceException)
	FindAllByB(name string, page dto.Page) (*domain.UserResponses, *dto.Paging, exception.ServiceException)
	FindAllArchivedB(page dto.Page) (*domain.UserResponses, *dto.Paging, exception.ServiceException)
	UpdateB(id uint, update *domain.UserUpdate) (*domain.UserResponse, exception.ServiceException)
	DeleteB(id uint) (*domain.UserResponse, exception.ServiceException)
}
