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

type gormUserS struct {
	db    *gorm.DB
	roleS abstraction.RoleS
}

func NewGormUserS(appCtx config.AppContext) abstraction.UserS {
	return gormUserS{db: appCtx.GetGormDB(), roleS: NewGormRoleS(appCtx)}
}

func (s gormUserS) ExistByUsername(username string) bool {
	err := s.db.Unscoped().Where("username = ?", username).First(&domain.User{}).Error
	return err != gorm.ErrRecordNotFound
}

func (s gormUserS) Save(domain *domain.User) (*domain.User, exception.ServiceException) {
	owning, queriedErr := s.roleS.FindAllById(domain.RoleIds)
	isOwningSideNotExisted := queriedErr != nil || len(*owning) == 0
	if isOwningSideNotExisted {
		return nil, exception.NewServiceException(queriedErr, constant.OwningSideNotExistedF)
	}
	domain.Roles = *owning
	savedErr := s.db.Create(domain).Error
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SaveF)
	}
	return domain, nil
}

func (s gormUserS) FindById(id uint) (*domain.User, exception.ServiceException) {
	domain := &domain.User{}
	err := s.db.Preload("Roles").First(domain, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByIdNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByIdF)
	}
	return domain, nil
}

func (s gormUserS) FindByUsername(username string) (*domain.User, exception.ServiceException) {
	domain := &domain.User{}
	err := s.db.Preload("Roles.Privileges").Where("username = ?", username).First(domain).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewServiceException(err, constant.FindByNoContentF)
		}
		return nil, exception.NewServiceException(err, constant.FindByF)
	}
	return domain, nil
}

func (s gormUserS) FindAllById(ids []uint) (*domain.Users, exception.ServiceException) {
	domains := &domain.Users{}
	err := s.db.Preload("Roles").Where("id IN ?", ids).Find(domains).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.FindAllByIdF)
	}
	return domains, nil
}

func (s gormUserS) FindAll(page dto.Page) (*domain.Users, *dto.Paging, exception.ServiceException) {
	domains := &domain.Users{}
	paging := &dto.Paging{}
	err := s.db.Scopes(PageScope(domains, s.db, page, paging)).Preload("Roles").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormUserS) FindAllBy(username string, page dto.Page) (*domain.Users, *dto.Paging, exception.ServiceException) {
	domains := &domain.Users{}
	paging := &dto.Paging{}
	err := s.db.Scopes(ConditionPageScope(domains, map[string]string{"Username": username}, s.db, page, paging)).Preload("Roles").Where("username LIKE ?", "%"+username+"%").Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllByF)
	}
	return domains, paging, nil
}

func (s gormUserS) FindAllArchived(page dto.Page) (*domain.Users, *dto.Paging, exception.ServiceException) {
	domains := &domain.Users{}
	paging := &dto.Paging{}
	err := s.db.Unscoped().Preload("Roles").Where("deleted_at IS NOT NULL").Scopes(ArchivedPageScope(domains, s.db, page, paging)).Find(domains).Error
	if err != nil {
		return nil, nil, exception.NewServiceException(err, constant.FindAllArchivedF)
	}
	return domains, paging, nil
}

func (s gormUserS) Update(id uint, update *domain.UserUpdate) (*domain.User, exception.ServiceException) {
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

func (s gormUserS) Delete(id uint) (*domain.User, exception.ServiceException) {
	old, queriedErr := s.FindById(id)
	if queriedErr != nil {
		if queriedErr.GetFailed() == constant.FindByIdNoContentF {
			return nil, exception.NewServiceException(queriedErr, constant.NotExistedF)
		}
		return nil, queriedErr
	}
	err := s.db.Where("id = ?", id).Delete(&domain.User{}).Error
	if err != nil {
		return nil, exception.NewServiceException(err, constant.DeleteF)
	}
	return old, nil
}
