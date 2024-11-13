package config

import (
	"fmt"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"sync"

	"github.com/casbin/casbin/v2"
)

type securityConfig struct {
	Enforcer *casbin.Enforcer
}

var (
	secConfigLock     = &sync.Mutex{}
	secConfigInstance *securityConfig
)

var (
	ErrSecConfig = exception.NewServiceException(nil, constant.SecConfigF)
)

func NewSecConfig() *securityConfig {
	return &securityConfig{}
}

func (cfg securityConfig) SecConfig() *securityConfig {
	if secConfigInstance == nil {
		secConfigLock.Lock()
		defer secConfigLock.Unlock()
		if secConfigInstance == nil {
			enforcer, enforcerErr := casbin.NewEnforcer("model.conf", "policy.csv")
			if enforcerErr != nil {
				fmt.Println("Cause: " + enforcerErr.Error())
				panic(constant.SecConfigF.Message)
			}
			secConfigInstance = &securityConfig{Enforcer: enforcer}
		}
	}
	return secConfigInstance
}
