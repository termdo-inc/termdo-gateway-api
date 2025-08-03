package auth

import (
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthProxy(apiBase, prefix string) gin.HandlerFunc {
	apiURL, _ := url.Parse(apiBase)
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	return func(c *gin.Context) {
		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, prefix)
		c.Request.Host = apiURL.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
