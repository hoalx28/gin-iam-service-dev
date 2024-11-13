package config

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type appContext struct {
	gormDB         *gorm.DB
	ginEngine      *gin.Engine
	casbinEnforcer *casbin.Enforcer
}

type AppContext interface {
	GetGormDB() *gorm.DB
	GetGinEngine() *gin.Engine
	GetCasbinEnforcer() *casbin.Enforcer
}

var (
	ctxLock            = &sync.Mutex{}
	appContextInstance *appContext
)

func NewAppContext(gormDB *gorm.DB, ginEngine *gin.Engine, casbinEnforcer *casbin.Enforcer) *appContext {
	if appContextInstance == nil {
		ctxLock.Lock()
		defer ctxLock.Unlock()
		if appContextInstance == nil {
			appContextInstance = &appContext{gormDB: gormDB, ginEngine: ginEngine, casbinEnforcer: casbinEnforcer}
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

func (ctx *appContext) GetCasbinEnforcer() *casbin.Enforcer {
	return ctx.casbinEnforcer
}
