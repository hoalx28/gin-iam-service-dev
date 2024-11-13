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

type statusT struct {
	business abstraction.StatusB
	httpUtil util.HttpUtil
}

func NewStatusT(appCtx config.AppContext) statusT {
	business := business.NewGormStatusB(appCtx)
	return statusT{business: business, httpUtil: util.NewHttpUtil()}
}

func (t statusT) Save(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creation domain.StatusCreation
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

func (t statusT) FindById(appCtx config.AppContext) gin.HandlerFunc {
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

func (t statusT) FindAll(appCtx config.AppContext) gin.HandlerFunc {
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

func (t statusT) FindAllBy(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page dto.Page
		if parseErr := ctx.ShouldBindQuery(&page); parseErr != nil {
			t.httpUtil.DoErrorParseQuery(ctx, parseErr)
			return
		}
		content := ctx.DefaultQuery("content", "")
		queried, paging, queriedErr := t.business.FindAllByB(content, page)
		if queriedErr != nil {
			t.httpUtil.DoError(ctx, queriedErr)
			return
		}
		t.httpUtil.DoSuccessPaging(ctx, constant.FindAllBy, queried, *paging)
	}
}

func (t statusT) FindAllArchived(appCtx config.AppContext) gin.HandlerFunc {
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

func (t statusT) Update(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var update domain.StatusUpdate
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

func (t statusT) Delete(appCtx config.AppContext) gin.HandlerFunc {
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
