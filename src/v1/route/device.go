package route

import (
	"iam/src/v1/config"
	"iam/src/v1/transport"
)

type deviceRoute struct{}

func NewDeviceRoute() *deviceRoute {
	return &deviceRoute{}
}

func (r deviceRoute) Config(appCtx config.AppContext) {
	transport := transport.NewDeviceTransport(appCtx)
	engine := appCtx.GetGinEngine()
	v1 := engine.Group("/api/v1")
	{
		devices := v1.Group("/devices")
		{
			devices.POST("/", transport.Save(appCtx))
			devices.GET("/:id", transport.FindById(appCtx))
			devices.GET("/", transport.FindAll(appCtx))
			devices.GET("/search", transport.FindAllBy(appCtx))
			devices.GET("/archived", transport.FindAllArchived(appCtx))
			devices.PATCH("/:id", transport.Update(appCtx))
			devices.DELETE("/:id", transport.Delete(appCtx))
		}
	}
}
