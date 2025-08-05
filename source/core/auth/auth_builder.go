package auth

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
)

const AuthApiPrefix = "/auth"

func BuildRoutes(app *gin.Engine) {
	app.Any(AuthApiPrefix+"/*rest", AuthProxy(config.AuthApiURL))
}
