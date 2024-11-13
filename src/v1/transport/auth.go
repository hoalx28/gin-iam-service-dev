package transport

import (
	"iam/src/v1/business"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/model/dto"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type authTransport struct {
	business business.AuthBusiness
	util     transportUtil
}

func NewAuthTransport(appCtx config.AppContext) *authTransport {
	business := business.NewGormAuthBusiness(appCtx)
	return &authTransport{business: business, util: NewTransportUtil()}
}

func (t authTransport) SignUp(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.RegisterRequest
		if parseErr := ctx.ShouldBind(&request); parseErr != nil {
			t.util.DoParseBodyErrorResponse(ctx, parseErr)
			return
		}
		credential, signUpErr := t.business.SignUp(request)
		if signUpErr != nil {
			t.util.DoErrorResponse(ctx, signUpErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.SignUp, credential)
	}
}

func (t authTransport) SignIn(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.CredentialRequest
		if parseErr := ctx.ShouldBind(&request); parseErr != nil {
			t.util.DoParseBodyErrorResponse(ctx, parseErr)
			return
		}
		credential, signUpErr := t.business.SignIn(request)
		if signUpErr != nil {
			t.util.DoErrorResponse(ctx, signUpErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.SignIn, credential)
	}
}

func (t authTransport) Identity(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.util.DoGetHeaderErrorResponse(ctx, AUTHORIZATION)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		_, identityErr := t.business.Identity(accessToken)
		if identityErr != nil {
			t.util.DoErrorResponse(ctx, identityErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.VerifyIdentity, true)
	}
}

func (t authTransport) Me(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.util.DoGetHeaderErrorResponse(ctx, AUTHORIZATION)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		register, queriedErr := t.business.Me(accessToken)
		if queriedErr != nil {
			t.util.DoErrorResponse(ctx, queriedErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.RetrieveProfile, register)
	}
}

func (t authTransport) SignOut(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.util.DoGetHeaderErrorResponse(ctx, AUTHORIZATION)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		tokenId, signOutErr := t.business.SignOut(accessToken)
		if signOutErr != nil {
			t.util.DoErrorResponse(ctx, signOutErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.SignOut, tokenId)
	}
}

func (t authTransport) Refresh(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.util.DoGetHeaderErrorResponse(ctx, AUTHORIZATION)
			return
		}
		refreshTokenHeader := ctx.GetHeader(X_REFRESH_TOKEN)
		if lo.IsEmpty(refreshTokenHeader) {
			t.util.DoGetHeaderErrorResponse(ctx, X_REFRESH_TOKEN)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		refreshToken := strings.Replace(refreshTokenHeader, "Bearer ", "", 1)
		tokenId, refreshErr := t.business.Refresh(accessToken, refreshToken)
		if refreshErr != nil {
			t.util.DoErrorResponse(ctx, refreshErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.RefreshToken, tokenId)
	}
}
