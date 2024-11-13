package main

import (
	"iam/src/v1/config"
	"iam/src/v1/route"
	"os"

	"github.com/gin-gonic/gin"
)

func routeConfig(appCtx config.AppContext) {
	privilegeRoute := route.NewPrivilegeR()
	roleRoute := route.NewRoleR()
	userRoute := route.NewUserR()
	deviceRoute := route.NewDeviceR()
	statusRoute := route.NewStatusR()
	authRoute := route.NewAuthR()

	privilegeRoute.Config(appCtx)
	roleRoute.Config(appCtx)
	userRoute.Config(appCtx)
	deviceRoute.Config(appCtx)
	statusRoute.Config(appCtx)
	authRoute.Config(appCtx)
}

func main() {
	engine := gin.Default()
	appCfg := config.NewAppConfig()
	appCfg.EnvConfig()
	db := appCfg.DBConfig()
	appCtx := config.NewAppContext(db, engine)
	appCfg.RecoverConfig(appCtx)
	appCfg.CorsConfig(appCtx)
	routeConfig(appCtx)
	port := os.Getenv("PORT")
	engine.Run(":" + port)
}
