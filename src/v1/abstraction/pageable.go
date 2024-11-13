package abstraction

import (
	"iam/src/v1/domain/dto"

	"gorm.io/gorm"
)

type Pageable interface {
	PageScope(value interface{}, db *gorm.DB, page dto.Page, paging *dto.Paging) func(db *gorm.DB) *gorm.DB
	ArchivedPageScope(value interface{}, db *gorm.DB, page dto.Page, paging *dto.Paging) func(db *gorm.DB) *gorm.DB
	ConditionPageScope(value interface{}, cond map[string]string, db *gorm.DB, page dto.Page, paging *dto.Paging) func(db *gorm.DB) *gorm.DB
}
