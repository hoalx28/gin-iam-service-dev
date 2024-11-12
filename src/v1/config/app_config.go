package config

import (
	"fmt"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"sync"

	"github.com/joho/godotenv"
)

type appConfig struct{}

type AppConfig interface {
	EnvConfig()
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
