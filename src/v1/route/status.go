package route

import (
	"iam/src/v1/config"
	"iam/src/v1/middleware"
	"iam/src/v1/transport"
)

type statusR struct{}

func NewStatusR() statusR {
	return statusR{}
}

func (r statusR) Config(appCtx config.AppContext) {
	transport := transport.NewStatusT(appCtx)
	middleware := middleware.NewAuthenticatedM(appCtx)
	engine := appCtx.GetGinEngine()
	v1 := engine.Group("/api/v1")
	{
		statuses := v1.Group("/statuses")
		{
			statuses.POST("/", transport.Save(appCtx))
			statuses.GET("/:id", middleware.Authenticated(), transport.FindById(appCtx))
			statuses.GET("/", transport.FindAll(appCtx))
			statuses.GET("/search", transport.FindAllBy(appCtx))
			statuses.GET("/archived", transport.FindAllArchived(appCtx))
			statuses.PATCH("/:id", transport.Update(appCtx))
			statuses.DELETE("/:id", transport.Delete(appCtx))
		}
	}
}
