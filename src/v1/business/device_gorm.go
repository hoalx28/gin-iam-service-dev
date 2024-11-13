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

type gormDeviceB struct {
	storage abstraction.DeviceS
}

func NewGormDeviceB(appCtx config.AppContext) abstraction.DeviceB {
	storage := storage.NewGormDeviceS(appCtx)
	return gormDeviceB{storage: storage}
}

func (b gormDeviceB) SaveB(creation *domain.DeviceCreation) (*domain.DeviceResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByIpAddress(*creation.IpAddress)
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

func (b gormDeviceB) FindByIdB(id uint) (*domain.DeviceResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return &response, nil
}

func (b gormDeviceB) FindAllByIdB(ids []uint) (*domain.DeviceResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormDeviceB) FindAllB(page dto.Page) (*domain.DeviceResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormDeviceB) FindAllByB(userAgent string, page dto.Page) (*domain.DeviceResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(userAgent, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormDeviceB) FindAllArchivedB(page dto.Page) (*domain.DeviceResponses, *dto.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormDeviceB) UpdateB(id uint, update *domain.DeviceUpdate) (*domain.DeviceResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return &response, nil
}

func (b gormDeviceB) DeleteB(id uint) (*domain.DeviceResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return &response, nil
}
