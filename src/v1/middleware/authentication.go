package middleware

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/token"
	"iam/src/v1/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type authenticatedM struct {
	jwtAuthProvider abstraction.JwtAuthProvider[dto.AuthClaims, domain.UserResponse]
	httpUtil        util.HttpUtil
}

func NewAuthenticatedM(appCtx config.AppContext) authenticatedM {
	jwtAuthProvider := token.NewJWTAuthProvider(appCtx)
	return authenticatedM{jwtAuthProvider: jwtAuthProvider, httpUtil: util.NewHttpUtil()}
}

func (m authenticatedM) Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(util.AUTHORIZATION)
		if lo.IsEmpty(header) {
			m.httpUtil.DoErrorGetHeader(ctx, util.AUTHORIZATION)
			return
		}
		accessToken := strings.Replace(header, "Bearer ", "", 1)
		claims, verifiedErr := m.jwtAuthProvider.EnsureNotBadCredential(accessToken)
		if verifiedErr != nil {
			m.httpUtil.DoError(ctx, verifiedErr)
			return
		}
		username := claims.Subject
		userId := claims.Payload.UserId
		referId := claims.Payload.ReferId
		scope := claims.Payload.Scope
		sessionId := claims.ID
		ctx.Set("username", username)
		ctx.Set("userId", userId)
		ctx.Set("referId", referId)
		ctx.Set("scope", scope)
		ctx.Set("sessionId", sessionId)
		ctx.Next()
	}
}
