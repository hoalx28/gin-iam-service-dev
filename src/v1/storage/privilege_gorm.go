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

type gormPrivilegeS struct {
	db *gorm.DB
}

func NewGormPrivilegeS(appCtx config.AppContext) abstraction.PrivilegeS {
	return gormPrivilegeS{db: appCtx.GetGormDB()}
}

func (s gormPrivilegeS) ExistByName(name string) bool {
	err := s.db.Unscoped().Where("name = ?", name).First(&domain.Privilege{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormPrivilegeS) Save(domain *domain.Privilege) (*domain.Privilege, exception.ServiceException) {
	err := s.db.Create(domain).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.SaveF)
	}
	return domain, nil
}

func (s gormPrivilegeS) FindById(id uint) (*domain.Privilege, exception.ServiceException) {
	domain := &domain.Privilege{}
	err := s.db.First(domain, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return domain, nil
}

func (s gormPrivilegeS) FindAllById(ids []uint) (*domain.Privileges, exception.ServiceException) {
	domains := &domain.Privileges{}
	err := s.db.Where("id IN ?", ids).Find(domains).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return domains, nil
}

func (s gormPrivilegeS) FindAll(page dto.Page) (*domain.Privileges, *dto.Paging, exception.ServiceException) {
	domains := &domain.Privileges{}
	paging := &dto.Paging{}
	err := s.db.Scopes(PageScope(domains, s.db, page, paging)).Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormPrivilegeS) FindAllBy(name string, page dto.Page) (*domain.Privileges, *dto.Paging, exception.ServiceException) {
	domains := &domain.Privileges{}
	paging := &dto.Paging{}
	err := s.db.Scopes(ConditionPageScope(domains, map[string]string{"Name": name}, s.db, page, paging)).Where("name LIKE ?", "%"+name+"%").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormPrivilegeS) FindAllArchived(page dto.Page) (*domain.Privileges, *dto.Paging, exception.ServiceException) {
	domains := &domain.Privileges{}
	paging := &dto.Paging{}
	err := s.db.Unscoped().Where("deleted_at IS NOT NULL").Scopes(ArchivedPageScope(domains, s.db, page, paging)).Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return domains, paging, nil
}

func (s gormPrivilegeS) Update(id uint, update *domain.PrivilegeUpdate) (*domain.Privilege, exception.ServiceException) {
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

func (s gormPrivilegeS) Delete(id uint) (*domain.Privilege, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&domain.Privilege{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
