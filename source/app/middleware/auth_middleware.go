package middleware

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/constants"
	"termdo.com/gateway-api/source/app/helpers"
)

func RequestAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if helpers.IsBrowserClient(ctx) {
			if token, ok := helpers.ReadAuthCookie(ctx); ok {
				ctx.Request.Header.Set("Authorization", constants.TokenPrefix+token)
			}
		}
		ctx.Next()
	}
}
