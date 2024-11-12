package storage

import (
	"gorm.io/gorm"
)

func PageScope(value interface{}, db *gorm.DB, page *Page, paging *Paging) func(db *gorm.DB) *gorm.DB {
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

func PageArchivedScope(value interface{}, db *gorm.DB, page *Page, paging *Paging) func(db *gorm.DB) *gorm.DB {
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

// ! Change to map condition filter
func PageNameFilterScope(value interface{}, name string, db *gorm.DB, page *Page, paging *Paging) func(db *gorm.DB) *gorm.DB {
	var totalRecord int64
	db.Model(value).Where("name LIKE ?", "%"+name+"%").Count(&totalRecord)
	totalPage := paging.TotalRecord/page.Size + 1
	paging.TotalRecord = int(totalRecord)
	paging.TotalPage = totalPage
	paging.Page = page.Page
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page.GetOffSet()).Limit(page.Size)
	}
}

func PageUsernameFilterScope(value interface{}, name string, db *gorm.DB, page *Page, paging *Paging) func(db *gorm.DB) *gorm.DB {
	var totalRecord int64
	db.Model(value).Where("username LIKE ?", "%"+name+"%").Count(&totalRecord)
	totalPage := paging.TotalRecord/page.Size + 1
	paging.TotalRecord = int(totalRecord)
	paging.TotalPage = totalPage
	paging.Page = page.Page
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page.GetOffSet()).Limit(page.Size)
	}
}

func PageUserAgentFilterScope(value interface{}, userAgent string, db *gorm.DB, page *Page, paging *Paging) func(db *gorm.DB) *gorm.DB {
	var totalRecord int64
	db.Model(value).Where("user_agent LIKE ?", "%"+userAgent+"%").Count(&totalRecord)
	totalPage := paging.TotalRecord/page.Size + 1
	paging.TotalRecord = int(totalRecord)
	paging.TotalPage = totalPage
	paging.Page = page.Page
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page.GetOffSet()).Limit(page.Size)
	}
}

func PageContentFilterScope(value interface{}, content string, db *gorm.DB, page *Page, paging *Paging) func(db *gorm.DB) *gorm.DB {
	var totalRecord int64
	db.Model(value).Where("content LIKE ?", "%"+content+"%").Count(&totalRecord)
	totalPage := paging.TotalRecord/page.Size + 1
	paging.TotalRecord = int(totalRecord)
	paging.TotalPage = totalPage
	paging.Page = page.Page
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page.GetOffSet()).Limit(page.Size)
	}
}
