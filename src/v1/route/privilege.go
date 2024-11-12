package route

import (
	"iam/src/v1/config"
	"iam/src/v1/transport"
)

type privilegeRoute struct{}

func NewPrivilegeRoute() *privilegeRoute {
	return &privilegeRoute{}
}

func (r privilegeRoute) Config(appCtx config.AppContext) {
	transport := transport.NewPrivilegeTransport(appCtx)
	engine := appCtx.GetGinEngine()
	v1 := engine.Group("/api/v1")
	{
		privileges := v1.Group("/privileges")
		{
			privileges.POST("/", transport.Save(appCtx))
			privileges.GET("/:id", transport.FindById(appCtx))
			privileges.GET("/", transport.FindAll(appCtx))
			privileges.GET("/search", transport.FindAllBy(appCtx))
			privileges.GET("/archived", transport.FindAllArchived(appCtx))
			privileges.PATCH("/:id", transport.Update(appCtx))
			privileges.DELETE("/:id", transport.Delete(appCtx))
		}
	}
}
