package route

import (
	"iam/src/v1/config"
	"iam/src/v1/transport"
)

type authR struct{}

func NewAuthR() *authR {
	return &authR{}
}

func (r authR) Config(appCtx config.AppContext) {
	transport := transport.NewAuthT(appCtx)
	engine := appCtx.GetGinEngine()
	v1 := engine.Group("/api/v1")
	{
		users := v1.Group("/auth")
		{
			users.POST("/sign-up", transport.SignUp(appCtx))
			users.POST("/sign-in", transport.SignIn(appCtx))
			users.POST("/identity", transport.Identity(appCtx))
			users.GET("/me", transport.Me(appCtx))
			users.POST("/sign-out", transport.SignOut(appCtx))
			users.POST("/refresh", transport.Refresh(appCtx))
		}
	}
}
