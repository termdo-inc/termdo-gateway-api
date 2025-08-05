package auth

import (
	"bytes"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/constants"
	"termdo.com/gateway-api/source/app/helpers"
	"termdo.com/gateway-api/source/app/utils"
)

func AuthProxy(apiBase string) gin.HandlerFunc {
	apiURL, _ := url.Parse(apiBase)
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	return func(ctx *gin.Context) {
		ctx.Request.URL.Path = strings.TrimPrefix(
			ctx.Request.URL.Path,
			AuthApiPrefix,
		)
		ctx.Request.Host = apiURL.Host

		rescp := &utils.ResponseCapture{
			ResponseWriter: ctx.Writer,
			Buffer:         &bytes.Buffer{},
		}
		proxy.ServeHTTP(rescp, ctx.Request)

		authApiHostname := rescp.Header().Get(constants.HeaderHostnameKey)
		rescp.Header().Del(constants.HeaderHostnameKey)

		if rescp.Buffer.Len() > 0 {
			helpers.SetHostnames(rescp, ctx, &authApiHostname, nil)
		} else {
			ctx.Writer.WriteHeader(rescp.Status)
		}
	}
}
