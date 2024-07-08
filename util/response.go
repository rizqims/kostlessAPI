package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SingleRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendSingleResponse(ctx *gin.Context, message string, data any, code int) {
	ctx.JSON(http.StatusOK, SingleRes{
		Code:    code,
		Message: message,
		Data:    data,
	})
}