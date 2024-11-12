package client

import (
	"fmt"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormClient struct {
	db *gorm.DB
}

var (
	gormLock           = &sync.Mutex{}
	gormClientInstance *gormClient
)

var (
	ErrConnectGormDB = exception.NewServiceException(nil, constant.DBConfigF)
)

func NewGormClient() *gormClient {
	return &gormClient{}
}

func (clt gormClient) Connect(dns string, cfg *gorm.Config) *gormClient {
	if gormClientInstance == nil {
		gormLock.Lock()
		defer gormLock.Unlock()
		if gormClientInstance == nil {
			db, err := gorm.Open(mysql.Open(dns), cfg)
			if err != nil {
				fmt.Println("Cause: " + err.Error())
				panic(constant.DBConfigF.Message)
			}
			gormClientInstance = &gormClient{db: db}
		}
	}
	return gormClientInstance
}

func (clt gormClient) GetDB() *gorm.DB {
	return clt.db
}
