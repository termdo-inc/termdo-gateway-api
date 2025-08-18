package tasks

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
)

const RoutePrefix = "/tasks"

func BuildRoutes(app *gin.Engine) {
	group := app.Group(RoutePrefix)
	group.Any("/*rest", Proxy(config.TasksApiURL))
}
