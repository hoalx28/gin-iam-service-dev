package transport

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/business"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain/dto"
	"iam/src/v1/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type authT struct {
	business abstraction.AuthB
	httpUtil util.HttpUtil
}

func NewAuthT(appCtx config.AppContext) authT {
	business := business.NewGormAuthB(appCtx)
	return authT{business: business, httpUtil: util.NewHttpUtil()}
}

func (t authT) SignUp(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.RegisterRequest
		if parseErr := ctx.ShouldBind(&request); parseErr != nil {
			t.httpUtil.DoErrorParseBody(ctx, parseErr)
			return
		}
		credential, signUpErr := t.business.SignUp(request)
		if signUpErr != nil {
			t.httpUtil.DoError(ctx, signUpErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.SignUp, credential)
	}
}

func (t authT) SignIn(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.CredentialRequest
		if parseErr := ctx.ShouldBind(&request); parseErr != nil {
			t.httpUtil.DoErrorParseBody(ctx, parseErr)
			return
		}
		credential, signUpErr := t.business.SignIn(request)
		if signUpErr != nil {
			t.httpUtil.DoError(ctx, signUpErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.SignIn, credential)
	}
}

func (t authT) Identity(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(util.AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.httpUtil.DoErrorGetHeader(ctx, util.AUTHORIZATION)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		_, identityErr := t.business.Identity(accessToken)
		if identityErr != nil {
			t.httpUtil.DoError(ctx, identityErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.VerifyIdentity, true)
	}
}

func (t authT) Me(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(util.AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.httpUtil.DoErrorGetHeader(ctx, util.AUTHORIZATION)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		register, queriedErr := t.business.Me(accessToken)
		if queriedErr != nil {
			t.httpUtil.DoError(ctx, queriedErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.RetrieveProfile, register)
	}
}

func (t authT) SignOut(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(util.AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.httpUtil.DoErrorGetHeader(ctx, util.AUTHORIZATION)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		tokenId, signOutErr := t.business.SignOut(accessToken)
		if signOutErr != nil {
			t.httpUtil.DoError(ctx, signOutErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.SignOut, tokenId)
	}
}

func (t authT) Refresh(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(util.AUTHORIZATION)
		if lo.IsEmpty(authorizationHeader) {
			t.httpUtil.DoErrorGetHeader(ctx, util.AUTHORIZATION)
			return
		}
		refreshTokenHeader := ctx.GetHeader(util.X_REFRESH_TOKEN)
		if lo.IsEmpty(refreshTokenHeader) {
			t.httpUtil.DoErrorGetHeader(ctx, util.X_REFRESH_TOKEN)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		refreshToken := strings.Replace(refreshTokenHeader, "Bearer ", "", 1)
		tokenId, refreshErr := t.business.Refresh(accessToken, refreshToken)
		if refreshErr != nil {
			t.httpUtil.DoError(ctx, refreshErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.RefreshToken, tokenId)
	}
}
