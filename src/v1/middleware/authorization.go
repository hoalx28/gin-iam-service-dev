package middleware

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
	"iam/src/v1/token"
	"iam/src/v1/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type authorizationM struct {
	jwtAuthProvider abstraction.JwtAuthProvider[dto.AuthClaims, domain.UserResponse]
	httpUtil        util.HttpUtil
}

func NewAuthorizationM(appCtx config.AppContext) authorizationM {
	jwtAuthProvider := token.NewJWTAuthProvider(appCtx)
	return authorizationM{jwtAuthProvider: jwtAuthProvider, httpUtil: util.NewHttpUtil()}
}

func getScope(ctx *gin.Context) (*string, exception.ServiceException) {
	rawScope, isExisted := ctx.Get("scope")
	if !isExisted {
		return nil, exception.NewServiceException(nil, constant.MissingAuthorizationHeaderF)
	}
	scope, ok := rawScope.(string)
	if !ok {
		return nil, exception.NewServiceException(nil, constant.IllLegalJwtTokenF)
	}
	return &scope, nil
}

func (m authorizationM) HasAuthority(authority string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		scope, scopeErr := getScope(ctx)
		if scopeErr != nil {
			m.httpUtil.DoError(ctx, scopeErr)
			return
		}
		if hasAuthority := strings.Contains(*scope, authority); !hasAuthority {
			m.httpUtil.DoErrorWith(ctx, constant.ForbiddenF)
			return
		}
		ctx.Next()
	}
}

func (m authorizationM) HasAnyAuthority(authorities []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isAllow := false
		scope, scopeErr := getScope(ctx)
		if scopeErr != nil {
			m.httpUtil.DoError(ctx, scopeErr)
			return
		}
		lo.ForEach(authorities, func(authority string, index int) {
			if hasAuthority := strings.Contains(*scope, authority); hasAuthority {
				isAllow = true
			}
		})
		if !isAllow {
			m.httpUtil.DoErrorWith(ctx, constant.ForbiddenF)
			return
		}
		ctx.Next()
	}
}
