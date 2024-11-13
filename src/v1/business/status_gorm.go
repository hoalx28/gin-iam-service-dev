package business

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
	"iam/src/v1/storage"
)

type gormStatusB struct {
	storage abstraction.StatusS
}

func NewGormStatusB(appCtx config.AppContext) abstraction.StatusB {
	storage := storage.NewGormStatusS(appCtx)
	return gormStatusB{storage: storage}
}

func (b gormStatusB) SaveB(creation *domain.StatusCreation) (*domain.StatusResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByContent(*creation.Content)
	if isExisted {
		return nil, exception.NewServiceException(nil, constant.AlreadyExistedF)
	}
	model := creation.AsModel()
	saved, saveErr := b.storage.Save(&model)
	if saveErr != nil {
		return nil, saveErr
	}
	response := saved.AsResponse()
	return &response, nil
}

func (b gormStatusB) FindByIdB(id uint) (*domain.StatusResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return &response, nil
}

func (b gormStatusB) FindAllByIdB(ids []uint) (*domain.StatusResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormStatusB) FindAllB(page dto.Page) (*domain.StatusResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormStatusB) FindAllByB(content string, page dto.Page) (*domain.StatusResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(content, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormStatusB) FindAllArchivedB(page dto.Page) (*domain.StatusResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormStatusB) UpdateB(id uint, update *domain.StatusUpdate) (*domain.StatusResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return &response, nil
}

func (b gormStatusB) DeleteB(id uint) (*domain.StatusResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return &response, nil
}
