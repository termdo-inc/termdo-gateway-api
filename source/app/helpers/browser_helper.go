package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
	"termdo.com/gateway-api/source/app/constants"
)

func IsBrowserClient(ctx *gin.Context) bool {
	return strings.EqualFold(
		ctx.GetHeader(constants.HeaderClientBrowserKey),
		"1",
	)
}

func ReadAuthCookie(ctx *gin.Context) (string, bool) {
	val, err := ctx.Cookie(constants.CookieName)
	if err != nil {
		return "", false
	}
	return val, true
}

func SetAuthCookie(ctx *gin.Context, token string) {
	ctx.SetSameSite(constants.CookieSameSite)
	ctx.SetCookie(
		constants.CookieName,
		token,
		0,
		constants.CookiePath,
		"",
		config.CookieSecure,
		true,
	)
}

func ClearAuthCookie(ctx *gin.Context) {
	ctx.SetSameSite(constants.CookieSameSite)
	ctx.SetCookie(
		constants.CookieName,
		"",
		-1,
		constants.CookiePath,
		"",
		config.CookieSecure,
		true,
	)
}
