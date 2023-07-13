package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TrailingSlash() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.URL.Path) > 1 && c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
			newPath := strings.TrimSuffix(c.Request.URL.Path, "/")
			c.Redirect(http.StatusMovedPermanently, newPath)
			return
		}
		c.Next()
	}
}
