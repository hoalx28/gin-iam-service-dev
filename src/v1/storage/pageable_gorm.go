package storage

import (
	"iam/src/v1/domain/dto"

	"gorm.io/gorm"
)

func PageScope(value interface{}, db *gorm.DB, page dto.Page, paging *dto.Paging) func(db *gorm.DB) *gorm.DB {
	var totalRecord int64
	db.Model(value).Count(&totalRecord)
	totalPage := paging.TotalRecord/page.Size + 1
	paging.TotalRecord = int(totalRecord)
	paging.TotalPage = totalPage
	paging.Page = page.Page
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page.GetOffSet()).Limit(page.Size)
	}
}

func ArchivedPageScope(value interface{}, db *gorm.DB, page dto.Page, paging *dto.Paging) func(db *gorm.DB) *gorm.DB {
	var totalRecord int64
	db.Model(value).Where("deleted_at IS NOT NULL").Count(&totalRecord)
	totalPage := paging.TotalRecord/page.Size + 1
	paging.TotalRecord = int(totalRecord)
	paging.TotalPage = totalPage
	paging.Page = page.Page
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page.GetOffSet()).Limit(page.Size)
	}
}

func ConditionPageScope(value interface{}, cond map[string]string, db *gorm.DB, page dto.Page, paging *dto.Paging) func(db *gorm.DB) *gorm.DB {
	var totalRecord int64
	db.Model(value).Where(cond).Count(&totalRecord)
	totalPage := paging.TotalRecord/page.Size + 1
	paging.TotalRecord = int(totalRecord)
	paging.TotalPage = totalPage
	paging.Page = page.Page
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page.GetOffSet()).Limit(page.Size)
	}
}
