package transport

import (
	"iam/src/v1/business"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/model"
	"iam/src/v1/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type privilegeTransport struct {
	business business.PrivilegeBusiness
	util     transportUtil
}

func NewPrivilegeTransport(appCtx config.AppContext) *privilegeTransport {
	business := business.NewGormPrivilegeBusiness(appCtx)
	return &privilegeTransport{business: business, util: NewTransportUtil()}
}

func (t privilegeTransport) Save(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creation model.PrivilegeCreation
		if parseErr := ctx.ShouldBind(&creation); parseErr != nil {
			t.util.DoParseBodyErrorResponse(ctx, parseErr)
			return
		}
		saved, savedErr := t.business.SaveBusiness(&creation)
		if savedErr != nil {
			t.util.DoErrorResponse(ctx, savedErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.Save, saved)
	}
}

func (t privilegeTransport) FindById(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, parseErr := strconv.Atoi(ctx.Param("id"))
		if parseErr != nil {
			t.util.DoParsePathErrorResponse(ctx, "id")
			return
		}
		queried, queriedErr := t.business.FindByIdBusiness(uint(id))
		if queriedErr != nil {
			t.util.DoErrorResponse(ctx, queriedErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.FindById, queried)
	}
}

func (t privilegeTransport) FindAll(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page storage.Page
		if parseErr := ctx.ShouldBindQuery(&page); parseErr != nil {
			t.util.DoParseQueryErrorResponse(ctx, parseErr)
			return
		}
		queried, paging, queriedErr := t.business.FindAllBusiness(&page)
		if queriedErr != nil {
			t.util.DoErrorResponse(ctx, queriedErr)
			return
		}
		t.util.DoSuccessPagingResponse(ctx, constant.FindAll, queried, *paging)
	}
}

func (t privilegeTransport) FindAllBy(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page storage.Page
		if parseErr := ctx.ShouldBindQuery(&page); parseErr != nil {
			t.util.DoParseQueryErrorResponse(ctx, parseErr)
			return
		}
		name := ctx.DefaultQuery("name", "")
		queried, paging, queriedErr := t.business.FindAllByBusiness(name, &page)
		if queriedErr != nil {
			t.util.DoErrorResponse(ctx, queriedErr)
			return
		}
		t.util.DoSuccessPagingResponse(ctx, constant.FindAllBy, queried, *paging)
	}
}

func (t privilegeTransport) FindAllArchived(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page storage.Page
		if parseErr := ctx.ShouldBindQuery(&page); parseErr != nil {
			t.util.DoParseQueryErrorResponse(ctx, parseErr)
			return
		}
		queried, paging, queriedErr := t.business.FindAllArchivedBusiness(&page)
		if queriedErr != nil {
			t.util.DoErrorResponse(ctx, queriedErr)
			return
		}
		t.util.DoSuccessPagingResponse(ctx, constant.FindAll, queried, *paging)
	}
}

func (t privilegeTransport) Update(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var update model.PrivilegeUpdate
		id, parseErr := strconv.Atoi(ctx.Param("id"))
		if parseErr != nil {
			t.util.DoParsePathErrorResponse(ctx, "id")
			return
		}
		if parseErr := ctx.ShouldBind(&update); parseErr != nil {
			t.util.DoParseBodyErrorResponse(ctx, parseErr)
			return
		}
		old, updateErr := t.business.UpdateBusiness(uint(id), &update)
		if updateErr != nil {
			t.util.DoErrorResponse(ctx, updateErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.Update, old)
	}
}

func (t privilegeTransport) Delete(appCtx config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, parseErr := strconv.Atoi(ctx.Param("id"))
		if parseErr != nil {
			t.util.DoParsePathErrorResponse(ctx, "id")
			return
		}
		old, deleteErr := t.business.DeleteBusiness(uint(id))
		if deleteErr != nil {
			t.util.DoErrorResponse(ctx, deleteErr)
			return
		}
		t.util.DoSuccessResponse(ctx, constant.Delete, old)
	}
}
