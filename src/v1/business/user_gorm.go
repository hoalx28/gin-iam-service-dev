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

type gormUserB struct {
	storage abstraction.UserS
}

func NewGormUserB(appCtx config.AppContext) abstraction.UserB {
	storage := storage.NewGormUserS(appCtx)
	return gormUserB{storage: storage}
}

func (b gormUserB) SaveB(creation *domain.UserCreation) (*domain.UserResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByUsername(*creation.Username)
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

func (b gormUserB) FindByIdB(id uint) (*domain.UserResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return &response, nil
}

func (b gormUserB) FindByUsernameB(username string) (*domain.UserResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindByUsername(username)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return &response, nil
}

func (b gormUserB) FindAllByIdB(ids []uint) (*domain.UserResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormUserB) FindAllB(page dto.Page) (*domain.UserResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormUserB) FindAllByB(name string, page dto.Page) (*domain.UserResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(name, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormUserB) FindAllArchivedB(page dto.Page) (*domain.UserResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormUserB) UpdateB(id uint, update *domain.UserUpdate) (*domain.UserResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return &response, nil
}

func (b gormUserB) DeleteB(id uint) (*domain.UserResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return &response, nil
}
