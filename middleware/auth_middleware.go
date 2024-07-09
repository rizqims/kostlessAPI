package middleware

import (
	"fmt"
	"kostless-api/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	CheckToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwt util.JwtToken
}

func (a *authMiddleware) CheckToken() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header require"})
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		fmt.Print("err ====", tokenString)
		claims, err := a.jwt.ValidateToken(tokenString)
		if err != nil {
			fmt.Print("error =====", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		ctx.Set("username", claims["username"])
		ctx.Next()
	}
}

func NewAuthMiddleware(jwtAuth util.JwtToken) AuthMiddleware {
	return &authMiddleware{jwt: jwtAuth}
}