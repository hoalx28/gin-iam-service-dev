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

type gormPrivilegeB struct {
	storage abstraction.PrivilegeS
}

func NewGormPrivilegeB(appCtx config.AppContext) abstraction.PrivilegeB {
	storage := storage.NewGormPrivilegeS(appCtx)
	return gormPrivilegeB{storage: storage}
}

func (b gormPrivilegeB) SaveB(creation *domain.PrivilegeCreation) (*domain.PrivilegeResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByName(*creation.Name)
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

func (b gormPrivilegeB) FindByIdB(id uint) (*domain.PrivilegeResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return &response, nil
}

func (b gormPrivilegeB) FindAllByIdB(ids []uint) (*domain.PrivilegeResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormPrivilegeB) FindAllB(page dto.Page) (*domain.PrivilegeResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormPrivilegeB) FindAllByB(name string, page dto.Page) (*domain.PrivilegeResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(name, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormPrivilegeB) FindAllArchivedB(page dto.Page) (*domain.PrivilegeResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormPrivilegeB) UpdateB(id uint, update *domain.PrivilegeUpdate) (*domain.PrivilegeResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return &response, nil
}

func (b gormPrivilegeB) DeleteB(id uint) (*domain.PrivilegeResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return &response, nil
}
