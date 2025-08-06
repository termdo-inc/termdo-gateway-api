package tasks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/config"
	"termdo.com/gateway-api/source/app/constants"
	"termdo.com/gateway-api/source/app/helpers"
	"termdo.com/gateway-api/source/app/utils"
	"termdo.com/gateway-api/source/core/auth/schemas"
)

func TasksProxy(apiBase string) gin.HandlerFunc {
	apiURL, _ := url.Parse(apiBase)
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		req, _ := http.NewRequest("GET", config.AuthApiURL+"/refresh", nil)
		req.Header.Set("Authorization", authHeader)
		res, err := http.DefaultClient.Do(req)
		if err != nil || res.StatusCode >= 400 {
			ctx.Status(res.StatusCode)
			body, _ := io.ReadAll(res.Body)
			ctx.Writer.Write(body)
			return
		}

		authApiHostname := res.Header.Get(constants.HeaderHostnameKey)
		res.Header.Del(constants.HeaderHostnameKey)

		var parsed schemas.RefreshResponse
		refreshBody, _ := io.ReadAll(res.Body)
		_ = json.Unmarshal(refreshBody, &parsed)

		ctx.Request.URL.Path = "/" + strconv.Itoa(parsed.Data.AccountID) +
			strings.TrimPrefix(ctx.Request.URL.Path, TasksApiPrefix)

		rescp := &utils.ResponseCapture{
			ResponseWriter: ctx.Writer,
			Buffer:         &bytes.Buffer{},
		}
		proxy.ServeHTTP(rescp, ctx.Request)

		tasksApiHostname := rescp.Header().Get(constants.HeaderHostnameKey)
		rescp.Header().Del(constants.HeaderHostnameKey)

		if rescp.Buffer.Len() > 0 {
			helpers.SetHostnames(rescp, ctx, &authApiHostname, &tasksApiHostname)

			var body map[string]any
			_ = json.Unmarshal(rescp.Buffer.Bytes(), &body)
			body["token"] = parsed.Token
			newBody, _ := json.Marshal(body)

			utils.CopyHeaders(rescp.Header(), ctx.Writer)

			ctx.Writer.Header().Set("Content-Length", strconv.Itoa(len(newBody)))
			ctx.Writer.WriteHeader(rescp.Status)
			ctx.Writer.Write(newBody)
		} else {
			ctx.Writer.WriteHeader(rescp.Status)
		}
	}
}
