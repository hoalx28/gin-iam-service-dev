package route

import (
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/middleware"
	"iam/src/v1/transport"
)

type statusR struct{}

func NewStatusR() statusR {
	return statusR{}
}

func (r statusR) Config(appCtx config.AppContext) {
	transport := transport.NewStatusT(appCtx)
	authenticatedM := middleware.NewAuthenticatedM(appCtx)
	authorizationM := middleware.NewAuthorizationM(appCtx)
	engine := appCtx.GetGinEngine()
	v1 := engine.Group("/api/v1")
	{
		statuses := v1.Group("/statuses")
		statuses.Use(authenticatedM.Authenticated())
		{
			statuses.POST("/", authorizationM.HasAuthority(constant.ROLE_ADMIN), transport.Save(appCtx))
			statuses.GET("/:id", authorizationM.HasAuthority(constant.ROLE_USER), transport.FindById(appCtx))
			statuses.GET("/", authorizationM.HasAuthority(constant.ROLE_USER), transport.FindAll(appCtx))
			statuses.GET("/search", authorizationM.HasAuthority(constant.ROLE_USER), transport.FindAllBy(appCtx))
			statuses.GET("/archived", authorizationM.HasAuthority(constant.ROLE_ADMIN), transport.FindAllArchived(appCtx))
			statuses.PATCH("/:id", authorizationM.HasAuthority(constant.ROLE_ADMIN), transport.Update(appCtx))
			statuses.DELETE("/:id", authorizationM.HasAuthority(constant.ROLE_ADMIN), transport.Delete(appCtx))
		}
	}
}
