package util

import (
	"errors"
	"iam/src/v1/constant"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
	"iam/src/v1/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	AUTHORIZATION   = "Authorization"
	X_REFRESH_TOKEN = "X-REFRESH-TOKEN"
)

type httpUtil struct{}

type HttpUtil interface {
	DoSuccess(ctx *gin.Context, c constant.Success, payload interface{})
	DoSuccessPaging(ctx *gin.Context, c constant.Success, payload interface{}, page dto.Paging)
	DoError(ctx *gin.Context, e exception.ServiceException)
	DoErrorParseBody(ctx *gin.Context, e error)
	DoErrorParseQuery(ctx *gin.Context, e error)
	DoErrorGetPath(ctx *gin.Context, name string)
	DoErrorGetHeader(ctx *gin.Context, name string)
}

func NewHttpUtil() httpUtil {
	return httpUtil{}
}

func (u httpUtil) DoSuccess(ctx *gin.Context, c constant.Success, payload interface{}) {
	response := response.NewResponse(int(time.Now().Unix()), c.Code, c.StatusCode, c.Message, payload)
	ctx.JSON(c.StatusCode, response)
}

func (u httpUtil) DoSuccessPaging(ctx *gin.Context, c constant.Success, payload interface{}, page dto.Paging) {
	paging := response.NewPaging(page.Page, page.TotalPage, page.TotalRecord)
	response := response.NewPagingResponse(int(time.Now().Unix()), c.Code, c.StatusCode, c.Message, payload, paging)
	ctx.JSON(c.StatusCode, response)
}

func (u httpUtil) DoError(ctx *gin.Context, e exception.ServiceException) {
	failed := e.GetFailed()
	response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, failed.Message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func (u httpUtil) DoErrorParseBody(ctx *gin.Context, e error) {
	var ve validator.ValidationErrors
	failed := constant.RequestBodyNotReadableF
	if errors.As(e, &ve) {
		message := msgForTag(strings.ToLower(ve[0].Field()), ve[0].Tag(), ve[0].Param())
		response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, message, nil)
		ctx.JSON(failed.StatusCode, response)
		return
	}
	response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, failed.Message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func (u httpUtil) DoErrorParseQuery(ctx *gin.Context, e error) {
	var ve validator.ValidationErrors
	failed := constant.RequestQueryNotReadableF
	if errors.As(e, &ve) {
		message := msgForTag(strings.ToLower(ve[0].Field()), ve[0].Tag(), ve[0].Param())
		response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, message, nil)
		ctx.JSON(failed.StatusCode, response)
		return
	}
	response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, failed.Message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func (u httpUtil) DoErrorGetPath(ctx *gin.Context, path string) {
	failed := constant.RequestParamsNotReadableF
	message := strings.ToLower(path) + " is missing in path variable list."
	response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func (u httpUtil) DoErrorGetHeader(ctx *gin.Context, headerName string) {
	failed := constant.RequestHeaderNotReadableF
	message := strings.ToLower(headerName) + " is missing in header list."
	response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func msgForTag(field string, tag string, value string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "gte":
		return field + " must be great than or equal " + value
	case "lte":
		return field + " must be less than or equal " + value
	default:
		return field + " is not legal"
	}
}
