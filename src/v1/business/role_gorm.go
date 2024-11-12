package business

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
)

type gormRoleBusiness struct {
	storage storage.RoleStorage
}

func NewGormRoleBusiness(appCtx config.AppContext) RoleBusiness {
	storage := storage.NewGormRoleStorage(appCtx)
	return gormRoleBusiness{storage: storage}
}

func (b gormRoleBusiness) SaveBusiness(creation *model.RoleCreation) (*model.RoleResponse, exception.ServiceException) {
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

func (b gormRoleBusiness) FindByIdBusiness(id uint) (*model.RoleResponse, exception.ServiceException) {
	model, queriedErr := b.storage.FindById(id)
	if queriedErr != nil {
		return nil, queriedErr
	}
	response := model.AsResponse()
	return response, nil
}

func (b gormRoleBusiness) FindAllByIdBusiness(ids []uint) (*model.RoleResponses, exception.ServiceException) {
	models, queriedErr := b.storage.FindAllById(ids)
	if queriedErr != nil {
		return nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, nil
}

func (b gormRoleBusiness) FindAllBusiness(page *storage.Page) (*model.RoleResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAll(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormRoleBusiness) FindAllByBusiness(name string, page *storage.Page) (*model.RoleResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllBy(name, page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormRoleBusiness) FindAllArchivedBusiness(page *storage.Page) (*model.RoleResponses, *storage.Paging, exception.ServiceException) {
	models, paging, queriedErr := b.storage.FindAllArchived(page)
	if queriedErr != nil {
		return nil, nil, queriedErr
	}
	responses := models.AsCollectionResponse()
	return &responses, paging, nil
}

func (b gormRoleBusiness) UpdateBusiness(id uint, update *model.RoleUpdate) (*model.RoleResponse, exception.ServiceException) {
	old, updateErr := b.storage.Update(id, update)
	if updateErr != nil {
		return nil, updateErr
	}
	response := old.AsResponse()
	return response, nil
}

func (b gormRoleBusiness) DeleteBusiness(id uint) (*model.RoleResponse, exception.ServiceException) {
	old, deleteErr := b.storage.Delete(id)
	if deleteErr != nil {
		return nil, deleteErr
	}
	response := old.AsResponse()
	return response, nil
}
