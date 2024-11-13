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

type gormRoleB struct {
	storage abstraction.RoleS
}

func NewGormRoleB(appCtx config.AppContext) abstraction.RoleB {
	storage := storage.NewGormRoleS(appCtx)
	return gormRoleB{storage: storage}
}

func (b gormRoleB) SaveB(creation *domain.RoleCreation) (*domain.RoleResponse, exception.ServiceException) {
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

func (b gormRoleB) FindByIdB(id uint) (*domain.RoleResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return &response, nil
}

func (b gormRoleB) FindAllByIdB(ids []uint) (*domain.RoleResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormRoleB) FindAllB(page dto.Page) (*domain.RoleResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormRoleB) FindAllByB(name string, page dto.Page) (*domain.RoleResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(name, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormRoleB) FindAllArchivedB(page dto.Page) (*domain.RoleResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormRoleB) UpdateB(id uint, update *domain.RoleUpdate) (*domain.RoleResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return &response, nil
}

func (b gormRoleB) DeleteB(id uint) (*domain.RoleResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return &response, nil
}
