package storage

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
)

type UserStorage interface {
	ExistByUsername(username string) bool
	Save(model *model.User) (*model.User, exception.ServiceException)
	FindById(id uint) (*model.User, exception.ServiceException)
	FindByUsername(username string) (*model.User, exception.ServiceException)
	FindAllById(ids []uint) (*model.Users, exception.ServiceException)
	FindAll(page *Page) (*model.Users, *Paging, exception.ServiceException)
	FindAllBy(username string, page *Page) (*model.Users, *Paging, exception.ServiceException)
	FindAllArchived(page *Page) (*model.Users, *Paging, exception.ServiceException)
	Update(id uint, update *model.UserUpdate) (*model.User, exception.ServiceException)
	Delete(id uint) (*model.User, exception.ServiceException)
}
