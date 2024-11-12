package route

import (
	"iam/src/v1/config"
	"iam/src/v1/transport"
)

type userRoute struct{}

func NewUserRoute() *userRoute {
	return &userRoute{}
}

func (r userRoute) Config(appCtx config.AppContext) {
	transport := transport.NewUserTransport(appCtx)
	engine := appCtx.GetGinEngine()
	v1 := engine.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", transport.Save(appCtx))
			users.GET("/:id", transport.FindById(appCtx))
			users.GET("/", transport.FindAll(appCtx))
			users.GET("/search", transport.FindAllBy(appCtx))
			users.GET("/archived", transport.FindAllArchived(appCtx))
			users.PATCH("/:id", transport.Update(appCtx))
			users.DELETE("/:id", transport.Delete(appCtx))
		}
	}
}
