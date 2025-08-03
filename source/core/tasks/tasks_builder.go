package tasks

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
)

func BuildRoutes(app *gin.Engine) {
	app.Any("/tasks/*rest", TasksProxy(config.TasksApiURL, "/tasks"))
}
