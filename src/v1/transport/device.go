package transport

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/business"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deviceT struct {
	business abstraction.DeviceB
	httpUtil util.HttpUtil
}

func NewDeviceT(appCtx config.AppContext) deviceT {
	business := business.NewGormDeviceB(appCtx)
	return deviceT{business: business, httpUtil: util.NewHttpUtil()}
}

func (t deviceT) Save(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creation domain.DeviceCreation
		if parseErr := ctx.ShouldBind(&creation); parseErr != nil {
			t.httpUtil.DoErrorParseBody(ctx, parseErr)
			return
		}
		saved, savedErr := t.business.SaveB(&creation)
		if savedErr != nil {
			t.httpUtil.DoError(ctx, savedErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.Save, saved)
	}
}

func (t deviceT) FindById(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, parseErr := strconv.Atoi(ctx.Param("id"))
		if parseErr != nil {
			t.httpUtil.DoErrorGetPath(ctx, "id")
			return
		}
		queried, queriedErr := t.business.FindByIdB(uint(id))
		if queriedErr != nil {
			t.httpUtil.DoError(ctx, queriedErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.FindById, queried)
	}
}

func (t deviceT) FindAll(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page dto.Page
		if parseErr := ctx.ShouldBindQuery(&page); parseErr != nil {
			t.httpUtil.DoErrorParseQuery(ctx, parseErr)
			return
		}
		queried, paging, queriedErr := t.business.FindAllB(page)
		if queriedErr != nil {
			t.httpUtil.DoError(ctx, queriedErr)
			return
		}
		t.httpUtil.DoSuccessPaging(ctx, constant.FindAll, queried, *paging)
	}
}

func (t deviceT) FindAllBy(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page dto.Page
		if parseErr := ctx.ShouldBindQuery(&page); parseErr != nil {
			t.httpUtil.DoErrorParseQuery(ctx, parseErr)
			return
		}
		userAgent := ctx.DefaultQuery("userAgent", "")
		queried, paging, queriedErr := t.business.FindAllByB(userAgent, page)
		if queriedErr != nil {
			t.httpUtil.DoError(ctx, queriedErr)
			return
		}
		t.httpUtil.DoSuccessPaging(ctx, constant.FindAllBy, queried, *paging)
	}
}

func (t deviceT) FindAllArchived(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page dto.Page
		if parseErr := ctx.ShouldBindQuery(&page); parseErr != nil {
			t.httpUtil.DoErrorParseQuery(ctx, parseErr)
			return
		}
		queried, paging, queriedErr := t.business.FindAllArchivedB(page)
		if queriedErr != nil {
			t.httpUtil.DoError(ctx, queriedErr)
			return
		}
		t.httpUtil.DoSuccessPaging(ctx, constant.FindAll, queried, *paging)
	}
}

func (t deviceT) Update(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var update domain.DeviceUpdate
		id, parseErr := strconv.Atoi(ctx.Param("id"))
		if parseErr != nil {
			t.httpUtil.DoErrorGetPath(ctx, "id")
			return
		}
		if parseErr := ctx.ShouldBind(&update); parseErr != nil {
			t.httpUtil.DoErrorParseBody(ctx, parseErr)
			return
		}
		old, updateErr := t.business.UpdateB(uint(id), &update)
		if updateErr != nil {
			t.httpUtil.DoError(ctx, updateErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.Update, old)
	}
}

func (t deviceT) Delete(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, parseErr := strconv.Atoi(ctx.Param("id"))
		if parseErr != nil {
			t.httpUtil.DoErrorGetPath(ctx, "id")
			return
		}
		old, deleteErr := t.business.DeleteB(uint(id))
		if deleteErr != nil {
			t.httpUtil.DoError(ctx, deleteErr)
			return
		}
		t.httpUtil.DoSuccess(ctx, constant.Delete, old)
	}
}
