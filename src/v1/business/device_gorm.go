package business

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type gormDeviceBusiness struct {
	storage storage.DeviceStorage
}

func NewGormDeviceBusiness(appCtx config.AppContext) DeviceBusiness {
	storage := storage.NewGormDeviceStorage(appCtx)
	return gormDeviceBusiness{storage: storage}
}

func (b gormDeviceBusiness) SaveBusiness(creation *model.DeviceCreation) (*model.DeviceResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByIpAddress(*creation.IpAddress)
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

func (b gormDeviceBusiness) FindByIdBusiness(id uint) (*model.DeviceResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return response, nil
}

func (b gormDeviceBusiness) FindAllByIdBusiness(ids []uint) (*model.DeviceResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormDeviceBusiness) FindAllBusiness(page *storage.Page) (*model.DeviceResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormDeviceBusiness) FindAllByBusiness(name string, page *storage.Page) (*model.DeviceResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(name, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormDeviceBusiness) FindAllArchivedBusiness(page *storage.Page) (*model.DeviceResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormDeviceBusiness) UpdateBusiness(id uint, update *model.DeviceUpdate) (*model.DeviceResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return response, nil
}

func (b gormDeviceBusiness) DeleteBusiness(id uint) (*model.DeviceResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return response, nil
}
