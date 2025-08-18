package auth

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
)

const RoutePrefix = "/auth"

func BuildRoutes(app *gin.Engine) {
	group := app.Group(RoutePrefix)
	group.Any("/*rest", func(c *gin.Context) {
		if c.Request.Method == "PUT" && c.Request.URL.Path == "/logout" {
			PutLogout()(c)
			return
		}
		Proxy(config.AuthApiURL)(c)
	})
}
