package transport

import (
	"errors"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/response"
	"iam/src/v1/storage"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type transportUtil struct{}

type TransportUtil interface {
	DoSuccessResponse(ctx *gin.Context, c constant.Success, payload interface{})
	DoSuccessPagingResponse(ctx *gin.Context, c constant.Success, payload interface{}, page storage.Paging)
	DoErrorResponse(ctx *gin.Context, e exception.ServiceException)
	DoParseBodyErrorResponse(ctx *gin.Context, e error)
	DoParsePathErrorResponse(ctx *gin.Context, name string)
	DoParseQueryErrorResponse(ctx *gin.Context, e error)
}

func NewTransportUtil() transportUtil {
	return transportUtil{}
}

func (u transportUtil) DoSuccessResponse(ctx *gin.Context, c constant.Success, payload interface{}) {
	response := response.NewResponse(time.Now().Unix(), c.Code, c.StatusCode, c.Message, payload)
	ctx.JSON(c.StatusCode, response)
}

func (u transportUtil) DoSuccessPagingResponse(ctx *gin.Context, c constant.Success, payload interface{}, page storage.Paging) {
	paging := response.NewPaging(page.Page, page.TotalPage, page.TotalRecord)
	response := response.NewPagingResponse(time.Now().Unix(), c.Code, c.StatusCode, c.Message, payload, paging)
	ctx.JSON(c.StatusCode, response)
}

func (u transportUtil) DoErrorResponse(ctx *gin.Context, e exception.ServiceException) {
	failed := e.GetFailed()
	response := response.NewResponse[interface{}](time.Now().Unix(), failed.Code, failed.StatusCode, failed.Message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func (u transportUtil) DoParseBodyErrorResponse(ctx *gin.Context, e error) {
	var ve validator.ValidationErrors
	failed := constant.RequestBodyNotReadableF
	if errors.As(e, &ve) {
		message := msgForTag(strings.ToLower(ve[0].Field()), ve[0].Tag(), ve[0].Param())
		response := response.NewResponse[interface{}](time.Now().Unix(), failed.Code, failed.StatusCode, message, nil)
		ctx.JSON(failed.StatusCode, response)
		return
	}
	response := response.NewResponse[interface{}](time.Now().Unix(), failed.Code, failed.StatusCode, failed.Message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func (u transportUtil) DoParseQueryErrorResponse(ctx *gin.Context, e error) {
	var ve validator.ValidationErrors
	failed := constant.RequestQueryNotReadableF
	if errors.As(e, &ve) {
		message := msgForTag(strings.ToLower(ve[0].Field()), ve[0].Tag(), ve[0].Param())
		response := response.NewResponse[interface{}](time.Now().Unix(), failed.Code, failed.StatusCode, message, nil)
		ctx.JSON(failed.StatusCode, response)
		return
	}
	response := response.NewResponse[interface{}](time.Now().Unix(), failed.Code, failed.StatusCode, failed.Message, nil)
	ctx.JSON(failed.StatusCode, response)
}

func (u transportUtil) DoParsePathErrorResponse(ctx *gin.Context, paramName string) {
	failed := constant.RequestParamsNotReadableF
	message := paramName + " is missing in path variable."
	response := response.NewResponse[interface{}](time.Now().Unix(), failed.Code, failed.StatusCode, message, nil)
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
