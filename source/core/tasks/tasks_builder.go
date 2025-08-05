package tasks

import (
	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
)

const TasksApiPrefix = "/tasks"

func BuildRoutes(app *gin.Engine) {
	app.Any(TasksApiPrefix+"/*rest", TasksProxy(config.TasksApiURL))
}
