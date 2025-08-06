package helpers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
	"termdo.com/gateway-api/source/app/constants"
	"termdo.com/gateway-api/source/app/utils"
)

func SetHostnames(
	rescp *utils.ResponseCapture,
	ctx *gin.Context,
	authApiHostname *string,
	tasksApiHostname *string,
) {
	var body map[string]any

	if err := json.Unmarshal(rescp.Buffer.Bytes(), &body); err != nil || body == nil {
		body = make(map[string]any)
	}

	hostnames := map[string]string{
		utils.KebabToCamelCase(config.AppHost): config.AppHostname,
	}
	if authApiHostname != nil {
		hostnames[utils.KebabToCamelCase(config.AuthApiHost)] = *authApiHostname
	}
	if tasksApiHostname != nil {
		hostnames[utils.KebabToCamelCase(config.TasksApiHost)] = *tasksApiHostname
	}
	body[constants.BodyHostnamesKey] = hostnames
	newBody, _ := json.Marshal(body)

	rescp.Header().Del(constants.HeaderHostnameKey)

	rescp.Buffer.Reset()
	rescp.Buffer.Write(newBody)
}
