package middlemock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthMiddlewareMock struct {
	mock.Mock
}

func (a *AuthMiddlewareMock) CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}