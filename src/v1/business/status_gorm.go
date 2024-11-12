package business

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type gormStatusBusiness struct {
	storage storage.StatusStorage
}

func NewGormStatusBusiness(appCtx config.AppContext) StatusBusiness {
	storage := storage.NewGormStatusStorage(appCtx)
	return gormStatusBusiness{storage: storage}
}

func (b gormStatusBusiness) SaveBusiness(creation *model.StatusCreation) (*model.StatusResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByContent(*creation.Content)
	if isExisted {
		return nil, exception.NewServiceException(nil, constant.AlreadyExistedF)
	}
	model := creation.AsModel()
	saved, saveErr := b.storage.Save(model)
	if saveErr != nil {
		return nil, saveErr
	}
	response := saved.AsResponse()
	return response, nil
}

func (b gormStatusBusiness) FindByIdBusiness(id uint) (*model.StatusResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return response, nil
}

func (b gormStatusBusiness) FindAllByIdBusiness(ids []uint) (*model.StatusResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormStatusBusiness) FindAllBusiness(page *storage.Page) (*model.StatusResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormStatusBusiness) FindAllByBusiness(content string, page *storage.Page) (*model.StatusResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(content, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormStatusBusiness) FindAllArchivedBusiness(page *storage.Page) (*model.StatusResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormStatusBusiness) UpdateBusiness(id uint, update *model.StatusUpdate) (*model.StatusResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return response, nil
}

func (b gormStatusBusiness) DeleteBusiness(id uint) (*model.StatusResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return response, nil
}
