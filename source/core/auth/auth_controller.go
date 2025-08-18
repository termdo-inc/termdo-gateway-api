package auth

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
	"termdo.com/gateway-api/source/app/helpers"
	"termdo.com/gateway-api/source/app/utils"
	"termdo.com/gateway-api/source/core/auth/schemas"
)

func PutLogout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if helpers.IsBrowserClient(ctx) {
			helpers.ClearAuthCookie(ctx)
		}

		body := schemas.BaseResponse[*any]{
			HttpStatus: schemas.HttpStatus{
				Code:    http.StatusOK,
				Message: http.StatusText(http.StatusOK),
			},
			ServerError:  nil,
			ClientErrors: []schemas.ClientError{},
			Data:         nil,
		}

		hostnames := map[string]string{
			utils.KebabToCamelCase(config.AppHost): config.AppHostname,
		}
		body.Hostnames = &hostnames

		payload, _ := json.Marshal(body)
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Content-Length", strconv.Itoa(len(payload)))
		ctx.Status(http.StatusOK)
		ctx.Writer.Write(payload)
	}
}
