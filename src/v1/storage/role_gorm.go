package storage

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"

	"gorm.io/gorm"
)

type gormRoleS struct {
	db         *gorm.DB
	privilegeS abstraction.PrivilegeS
}

func NewGormRoleS(appCtx config.AppContext) abstraction.RoleS {
	return gormRoleS{db: appCtx.GetGormDB(), privilegeS: NewGormPrivilegeS(appCtx)}
}

func (s gormRoleS) ExistByName(name string) bool {
	err := s.db.Unscoped().Where("name = ?", name).First(&domain.Role{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormRoleS) Save(domain *domain.Role) (*domain.Role, exception.ServiceException) {
	owning, queriedErr := s.privilegeS.FindAllById(domain.PrivilegeIds)
	isOwningSideNotExisted := queriedErr != nil || len(*owning) == 0
	if isOwningSideNotExisted {
		return nil, exception.NewServiceException(queriedErr, constant.OwningSideNotExistedF)
	}
	domain.Privileges = *owning
	savedErr := s.db.Create(domain).Error
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SaveF)
	}
	return domain, nil
}

func (s gormRoleS) FindById(id uint) (*domain.Role, exception.ServiceException) {
	domain := &domain.Role{}
	err := s.db.Preload("Privileges").First(domain, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return domain, nil
}

func (s gormRoleS) FindAllById(ids []uint) (*domain.Roles, exception.ServiceException) {
	domains := &domain.Roles{}
	err := s.db.Preload("Privileges").Where("id IN ?", ids).Find(domains).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return domains, nil
}

func (s gormRoleS) FindAll(page dto.Page) (*domain.Roles, *dto.Paging, exception.ServiceException) {
	domains := &domain.Roles{}
	paging := &dto.Paging{}
	err := s.db.Scopes(PageScope(domains, s.db, page, paging)).Preload("Privileges").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormRoleS) FindAllBy(name string, page dto.Page) (*domain.Roles, *dto.Paging, exception.ServiceException) {
	domains := &domain.Roles{}
	paging := &dto.Paging{}
	err := s.db.Scopes(ConditionPageScope(domains, map[string]string{"Name": name}, s.db, page, paging)).Preload("Privileges").Where("name LIKE ?", "%"+name+"%").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormRoleS) FindAllArchived(page dto.Page) (*domain.Roles, *dto.Paging, exception.ServiceException) {
	domains := &domain.Roles{}
	paging := &dto.Paging{}
	err := s.db.Unscoped().Preload("Privileges").Where("deleted_at IS NOT NULL").Scopes(ArchivedPageScope(domains, s.db, page, paging)).Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return domains, paging, nil
}

func (s gormRoleS) Update(id uint, update *domain.RoleUpdate) (*domain.Role, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Updates(update).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.UpdateF)
	}
	return old, nil
}

func (s gormRoleS) Delete(id uint) (*domain.Role, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&domain.Role{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
