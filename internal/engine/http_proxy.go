package engine

import (
	"net/http/httputil"
	"net/url"

	"github.com/issueye/grape/internal/global"
)

// http 代理
type HttpProxy struct {
	IsTSL bool // 是否需要证书
	Port  uint // 端口号
}

// 定义反向代理处理器
func ReverseProxyHttpHandler(targetURL string) *httputil.ReverseProxy {
	// 解析目标URL
	target, err := url.Parse(targetURL)
	if err != nil {
		global.Log.Errorf("解析目标URL失败 %s", err.Error())
	}

	// 创建反向代理实例
	proxy := httputil.NewSingleHostReverseProxy(target)

	// return func(c *gin.Context) {
	// 	// 转发请求到目标服务器
	// 	// 处理url
	// 	path := c.Param("path")
	// 	fmt.Println("c.Request.URL.Path", path)
	// 	c.Request.URL.Path = path
	// 	proxy.ServeHTTP(c.Writer, c.Request)
	// }
	return proxy
}
