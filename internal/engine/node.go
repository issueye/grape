package engine

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

func LoadNode(portId string, engine *gin.Engine) {
	nodeList, err := logic.Node{}.Get(&repository.QueryNode{
		PortId: portId,
	})
	if err != nil {
		return
	}

	for _, nodeElem := range nodeList {
		if nodeElem.NodeType == 1 {
			node := nodeElem.Name
			target := nodeElem.Target
			proxy := ReverseProxyHttpHandler(target)
			rg := engine.Group(node)
			rg.Any("*path", func(ctx *gin.Context) {
				path := ctx.Param("path")
				fmt.Println("path", path, node)
				if path[:5] == "/web/" {
					if strings.Contains(path, "/web/isInitServer") {
						ctx.Request.URL.Path = "/isInitServer"
						proxy.ServeHTTP(ctx.Writer, ctx.Request)
						return
					}

					dir := fmt.Sprintf("./runtime/static/%s", node)
					fmt.Println("dir", dir)
					staticRouterPath := fmt.Sprintf("/%s/web/", node)
					fmt.Println("staticRouterPath", staticRouterPath)
					http.StripPrefix(staticRouterPath, http.FileServer(http.Dir(dir))).ServeHTTP(ctx.Writer, ctx.Request)
					return
				}

				// 处理节点名称
				ctx.Request.URL.Path = path
				proxy.ServeHTTP(ctx.Writer, ctx.Request)
			})

			// rg.Any("*path", ReverseProxyHandler(nodeElem.Target))
			// rg.Static("/web/", fmt.Sprintf("./runtime/static/%s", nodeElem.Name))
		} else {
			// engine.Any(fmt.Sprintf("%s/*path", nodeElem.Name), ReverseProxyHandler(nodeElem.Target))
		}
	}
}

// 定义反向代理处理器
func ReverseProxyHandler(targetURL string) gin.HandlerFunc {
	// 解析目标URL
	target, err := url.Parse(targetURL)
	if err != nil {
		global.Log.Errorf("解析目标URL失败 %s", err.Error())
	}

	// 创建反向代理实例
	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		// 转发请求到目标服务器
		// 处理url
		path := c.Param("path")
		fmt.Println("c.Request.URL.Path", path)
		c.Request.URL.Path = path
		proxy.ServeHTTP(c.Writer, c.Request)
	}
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
