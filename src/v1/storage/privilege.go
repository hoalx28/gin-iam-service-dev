package storage

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
)

type PrivilegeStorage interface {
	ExistByName(name string) bool
	Save(model *model.Privilege) (*model.Privilege, exception.ServiceException)
	FindById(id uint) (*model.Privilege, exception.ServiceException)
	FindAllById(ids []uint) (*model.Privileges, exception.ServiceException)
	FindAll(page *Page) (*model.Privileges, *Paging, exception.ServiceException)
	FindAllBy(name string, page *Page) (*model.Privileges, *Paging, exception.ServiceException)
	FindAllArchived(page *Page) (*model.Privileges, *Paging, exception.ServiceException)
	Update(id uint, update *model.PrivilegeUpdate) (*model.Privilege, exception.ServiceException)
	Delete(id uint) (*model.Privilege, exception.ServiceException)
}
