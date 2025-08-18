package auth

import (
	"bytes"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/constants"
	"termdo.com/gateway-api/source/app/helpers"
	"termdo.com/gateway-api/source/app/utils"
)

func Proxy(apiBase string) gin.HandlerFunc {
	apiURL, _ := url.Parse(apiBase)
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	return func(ctx *gin.Context) {
		ctx.Request.URL.Path = strings.TrimPrefix(
			ctx.Request.URL.Path,
			RoutePrefix,
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
			helpers.SetHostnames(ctx, rescp, &authApiHostname, nil)
			helpers.NormalizeTokenResponse(ctx, rescp, nil)

			newBody := rescp.Buffer.Bytes()

			utils.CopyHeaders(rescp.Header(), ctx.Writer)

			ctx.Writer.Header().Set("Content-Length", strconv.Itoa(len(newBody)))
			ctx.Writer.WriteHeader(rescp.Status)
			ctx.Writer.Write(newBody)
		} else {
			ctx.Writer.WriteHeader(rescp.Status)
		}
	}
}
