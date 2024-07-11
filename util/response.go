package util

import (
	"log"

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

func SendSingleResponse(c *gin.Context, code int, message string, data any) {
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
func SendErrRes(c *gin.Context, code int, message string) {
	c.JSON(code, errRes{
		Code:    code,
		Message: message,
	})
}

func SendEmail(to, subject, body string) error {
	log.Printf("Sending email to: %s, Subject: %s, Body: %s\n", to, subject, body)
	return nil
}

func NotifyOwner(message string) error {
	log.Printf("Notifying owner: %s\n", message)
	return nil
}
