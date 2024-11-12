package main

import (
	"iam/src/v1/client"
	"iam/src/v1/config"
	"iam/src/v1/model"
	"iam/src/v1/route"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type restServer struct{}

func (rest restServer) envConfig(appCfg config.AppConfig) {
	appCfg.EnvConfig()
}

func (rest restServer) dbConfig(dns string) *gorm.DB {
	gormClt := client.NewGormClient()
	gormClt = gormClt.Connect(dns, &gorm.Config{})
	db := gormClt.GetDB()
	models := []interface{}{&model.Privilege{}, &model.Role{}, &model.User{}, &model.Device{}, &model.Status{}}
	db.AutoMigrate(models...)
	return db
}

func (rest restServer) routeConfig(appCtx config.AppContext) {
	privilegeRoute := route.NewPrivilegeRoute()
	roleRoute := route.NewRoleRoute()
	userRoute := route.NewUserRoute()
	deviceRoute := route.NewDeviceRoute()
	statusRoute := route.NewStatusRoute()
	privilegeRoute.Config(appCtx)
	roleRoute.Config(appCtx)
	userRoute.Config(appCtx)
	deviceRoute.Config(appCtx)
	statusRoute.Config(appCtx)
}

func (rest restServer) corsConfig(appCtx config.AppContext) {
	ginEngine := appCtx.GetGinEngine()
	ginEngine.Use(cors.Default())
}

func main() {
	appCfg := config.NewAppConfig()
	restService := restServer{}
	restService.envConfig(appCfg)
	dns := os.Getenv("GORM_DNS")
	port := os.Getenv("PORT")
	db := restService.dbConfig(dns)
	engine := gin.Default()
	appCtx := config.NewAppContext(db, engine)
	restService.routeConfig(appCtx)
	restService.corsConfig(appCtx)
	engine.Run(":" + port)
}