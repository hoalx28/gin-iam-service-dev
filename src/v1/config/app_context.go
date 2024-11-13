package config

import (
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type appContext struct {
	gormDB    *gorm.DB
	ginEngine *gin.Engine
}

type AppContext interface {
	GetGormDB() *gorm.DB
	GetGinEngine() *gin.Engine
}

var (
	ctxLock            = &sync.Mutex{}
	appContextInstance *appContext
)

func NewAppContext(gormDB *gorm.DB, ginEngine *gin.Engine) *appContext {
	if appContextInstance == nil {
		ctxLock.Lock()
		defer ctxLock.Unlock()
		if appContextInstance == nil {
			appContextInstance = &appContext{gormDB: gormDB, ginEngine: ginEngine}
		}
	}
	return appContextInstance
}

func (ctx *appContext) GetGormDB() *gorm.DB {
	return ctx.gormDB
}

func (ctx *appContext) GetGinEngine() *gin.Engine {
	return ctx.ginEngine
}
