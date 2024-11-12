package business

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type gormPrivilegeBusiness struct {
	storage storage.PrivilegeStorage
}

func NewGormPrivilegeBusiness(appCtx config.AppContext) PrivilegeBusiness {
	storage := storage.NewGormPrivilegeStorage(appCtx)
	return gormPrivilegeBusiness{storage: storage}
}

func (b gormPrivilegeBusiness) SaveBusiness(creation *model.PrivilegeCreation) (*model.PrivilegeResponse, exception.ServiceException) {
	isExisted := b.storage.ExistByName(*creation.Name)
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

func (b gormPrivilegeBusiness) FindByIdBusiness(id uint) (*model.PrivilegeResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return response, nil
}

func (b gormPrivilegeBusiness) FindAllByIdBusiness(ids []uint) (*model.PrivilegeResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormPrivilegeBusiness) FindAllBusiness(page *storage.Page) (*model.PrivilegeResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormPrivilegeBusiness) FindAllByBusiness(name string, page *storage.Page) (*model.PrivilegeResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(name, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormPrivilegeBusiness) FindAllArchivedBusiness(page *storage.Page) (*model.PrivilegeResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormPrivilegeBusiness) UpdateBusiness(id uint, update *model.PrivilegeUpdate) (*model.PrivilegeResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return response, nil
}

func (b gormPrivilegeBusiness) DeleteBusiness(id uint) (*model.PrivilegeResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return response, nil
}
