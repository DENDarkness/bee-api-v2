package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(token string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.GetHeader("Authorization")
		if t == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is empty"})
			return
		}
		if t != token {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
			return
		}
	}
}
