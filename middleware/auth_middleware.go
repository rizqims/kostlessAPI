package middleware

import (
	"kostless-api/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthoMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header require"})
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}