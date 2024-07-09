package util

import (
	"github.com/gin-gonic/gin"
)

type singleRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type errRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendSingleResponse(c *gin.Context,code int, message string, data any) {
	c.JSON(code, singleRes{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendErrResponse(c *gin.Context, code int, message string) {
	c.JSON(code, errRes{
		Code:    code,
		Message: message,
	})
}
