package config

import (
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/response"
	"time"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if exception, ok := err.(exception.ServiceException); ok {
					failed := exception.GetFailed()
					response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, failed.Message, nil)
					c.AbortWithStatusJSON(failed.StatusCode, response)
					panic(err)
				}
				failed := constant.UncategorizedF
				response := response.NewResponse[interface{}](int(time.Now().Unix()), failed.Code, failed.StatusCode, failed.Message, nil)
				c.AbortWithStatusJSON(failed.StatusCode, response)
				panic(err)
			}
		}()

		c.Next()
	}
}
