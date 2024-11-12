package route

import (
	"iam/src/v1/config"
	"iam/src/v1/transport"
)

type roleRoute struct{}

func NewRoleRoute() *roleRoute {
	return &roleRoute{}
}

func (r roleRoute) Config(appCtx config.AppContext) {
	transport := transport.NewRoleTransport(appCtx)
	engine := appCtx.GetGinEngine()
	v1 := engine.Group("/api/v1")
	{
		roles := v1.Group("/roles")
		{
			roles.POST("/", transport.Save(appCtx))
			roles.GET("/:id", transport.FindById(appCtx))
			roles.GET("/", transport.FindAll(appCtx))
			roles.GET("/search", transport.FindAllBy(appCtx))
			roles.GET("/archived", transport.FindAllArchived(appCtx))
			roles.PATCH("/:id", transport.Update(appCtx))
			roles.DELETE("/:id", transport.Delete(appCtx))
		}
	}
}
