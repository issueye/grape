package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var ruleMap = map[string]string{
	`/admin/api/`: `/api/`,
}

func Intercept(e *gin.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.Path
		for k, v := range ruleMap {
			if strings.Contains(url, k) {
				newPath := strings.ReplaceAll(url, k, v)
				ctx.Request.URL.Path = newPath
				e.HandleContext(ctx)
				ctx.Abort()
				break
			}
		}
		ctx.Next()
	}
}
