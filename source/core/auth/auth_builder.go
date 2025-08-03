package auth

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
)

func BuildRoutes(app *gin.Engine) {
	app.Any("/auth/*rest", AuthProxy(config.AuthApiURL, "/auth"))
}
