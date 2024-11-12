package storage

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
)

type RoleStorage interface {
	ExistByName(name string) bool
	Save(model *model.Role) (*model.Role, exception.ServiceException)
	FindById(id uint) (*model.Role, exception.ServiceException)
	FindAllById(ids []uint) (*model.Roles, exception.ServiceException)
	FindAll(page *Page) (*model.Roles, *Paging, exception.ServiceException)
	FindAllBy(name string, page *Page) (*model.Roles, *Paging, exception.ServiceException)
	FindAllArchived(page *Page) (*model.Roles, *Paging, exception.ServiceException)
	Update(id uint, update *model.RoleUpdate) (*model.Role, exception.ServiceException)
	Delete(id uint) (*model.Role, exception.ServiceException)
}
