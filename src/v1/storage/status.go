package storage

import (
	"iam/src/v1/exception"
	"iam/src/v1/model"
)

type StatusStorage interface {
	ExistByContent(content string) bool
	Save(model *model.Status) (*model.Status, exception.ServiceException)
	FindById(id uint) (*model.Status, exception.ServiceException)
	FindAllById(ids []uint) (*model.Statuses, exception.ServiceException)
	FindAll(page *Page) (*model.Statuses, *Paging, exception.ServiceException)
	FindAllBy(content string, page *Page) (*model.Statuses, *Paging, exception.ServiceException)
	FindAllArchived(page *Page) (*model.Statuses, *Paging, exception.ServiceException)
	Update(id uint, update *model.StatusUpdate) (*model.Status, exception.ServiceException)
	Delete(id uint) (*model.Status, exception.ServiceException)
}
