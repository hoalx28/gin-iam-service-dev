package business

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type gormUserBusiness struct {
	storage storage.UserStorage
}

func NewGormUserBusiness(appCtx config.AppContext) UserBusiness {
	storage := storage.NewGormUserStorage(appCtx)
	return gormUserBusiness{storage: storage}
}

func (b gormUserBusiness) SaveBusiness(creation *model.UserCreation) (*model.UserResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByUsername(*creation.Username)
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

func (b gormUserBusiness) FindByIdBusiness(id uint) (*model.UserResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return response, nil
}

func (b gormUserBusiness) FindByUsernameBusiness(username string) (*model.UserResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindByUsername(username)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return response, nil
}

func (b gormUserBusiness) FindAllByIdBusiness(ids []uint) (*model.UserResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormUserBusiness) FindAllBusiness(page *storage.Page) (*model.UserResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormUserBusiness) FindAllByBusiness(name string, page *storage.Page) (*model.UserResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(name, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormUserBusiness) FindAllArchivedBusiness(page *storage.Page) (*model.UserResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormUserBusiness) UpdateBusiness(id uint, update *model.UserUpdate) (*model.UserResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return response, nil
}

func (b gormUserBusiness) DeleteBusiness(id uint) (*model.UserResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return response, nil
}
