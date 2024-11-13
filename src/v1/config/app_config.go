package config

import (
	"fmt"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/exception"
	"os"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type appConfig struct{}

type AppConfig interface {
	EnvConfig()
	DBConfig(dns string) *gorm.DB
	RecoverConfig(appCtx AppContext)
	CorsConfig(appCtx AppContext)
}

var (
	cfgLock           = &sync.Mutex{}
	appConfigInstance *appConfig
)

var (
	ErrEnvConfig = exception.NewServiceException(nil, constant.EnvConfigF)
)

func NewAppConfig() *appConfig {
	if appConfigInstance == nil {
		cfgLock.Lock()
		defer cfgLock.Unlock()
		if appConfigInstance == nil {
			appConfigInstance = &appConfig{}
		}
	}
	return appConfigInstance
}

func (cfg appConfig) EnvConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Cause: " + err.Error())
		panic(constant.EnvConfigF.Message)
	}
}

func (cfg appConfig) DBConfig() *gorm.DB {
	dns := os.Getenv("GORM_DNS")
	gormClt := NewGormClient().Connect(dns, &gorm.Config{})
	db := gormClt.GetDB()
	domains := []interface{}{&domain.Privilege{}, &domain.Role{}, &domain.User{}, &domain.Device{}, &domain.Status{}, &domain.BadCredential{}}
	db.AutoMigrate(domains...)
	return db
}

func (cfg appConfig) RecoverConfig(appCtx AppContext) {
	ginEngine := appCtx.GetGinEngine()
	ginEngine.Use(Recover())
}

func (cfg appConfig) CorsConfig(appCtx AppContext) {
	ginEngine := appCtx.GetGinEngine()
	ginEngine.Use(cors.Default())
}
