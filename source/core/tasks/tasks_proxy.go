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
	"termdo.com/gateway-api/source/app/utils"
	"termdo.com/gateway-api/source/core/tasks/schemas"
)

func TasksProxy(apiBase, prefix string) gin.HandlerFunc {
	apiURL, _ := url.Parse(apiBase)
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization header",
			})
			return
		}

		req, _ := http.NewRequest("GET", config.AuthApiURL+"/refresh", nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil || res.StatusCode >= 400 {
			c.Status(res.StatusCode)
			body, _ := io.ReadAll(res.Body)
			c.Writer.Write(body)
			return
		}

		var parsed schemas.RefreshResponse
		body, _ := io.ReadAll(res.Body)
		_ = json.Unmarshal(body, &parsed)

		c.Request.URL.Path = "/" + strconv.Itoa(parsed.AccountID) +
			strings.TrimPrefix(c.Request.URL.Path, prefix)

		rec := &utils.ResponseCapture{
			ResponseWriter: c.Writer,
			Buffer:         &bytes.Buffer{},
		}
		proxy.ServeHTTP(rec, c.Request)

		c.Writer.WriteHeader(rec.Status)
		var proxyBody map[string]any
		_ = json.Unmarshal(rec.Buffer.Bytes(), &proxyBody)
		proxyBody["token"] = parsed.Token

		newBody, _ := json.Marshal(proxyBody)
		c.Writer.Write(newBody)
	}
}
