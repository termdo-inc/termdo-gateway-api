package auth

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
)

const RoutePrefix = "/auth"

func BuildRoutes(app *gin.Engine) {
	app.PUT(RoutePrefix+"/logout", PutLogout())
	app.Any(RoutePrefix+"/*rest", Proxy(config.AuthApiURL))
}
