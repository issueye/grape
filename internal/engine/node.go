package engine

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

func LoadNode(portId string, engine *gin.Engine) {
	nodeList, err := logic.Node{}.Get(&repository.QueryNode{
		PortId: portId,
	})
	if err != nil {
		return
	}

	for _, nodeElem := range nodeList {
		dir := filepath.Join("runtime", "static", "pages", nodeElem.PortId, nodeElem.Name, nodeElem.FileName)
		fmt.Println("静态文件路径：", dir)
		handler := http.FileServer(http.Dir(dir))
		node := nodeElem.Name
		target := nodeElem.Target
		// pagePath := nodeElem.PagePath
		proxy := ReverseProxyHttpHandler(target)
		rg := engine.Group(node)

		ruleList, err := service.NewRule().Query(&repository.QueryRule{
			PortId: portId,
			NodeId: nodeElem.ID,
		})
		if err != nil {
			continue
		}

		r, err := url.JoinPath(fmt.Sprintf("/%s", node), "web")
		if err != nil {
			global.Log.Errorf("路由拼接失败 %s", err.Error())
			continue
		}

		if nodeElem.NodeType == 1 {
			rg.Any("*path", func(ctx *gin.Context) {
				path := ctx.Param("path")
				fmt.Println("path", path)
				for _, nodeRule := range ruleList {
					fmt.Println("nodeRule   ", path, nodeRule.Name)
					re := regexp.MustCompile(nodeRule.Name)
					match := re.FindStringSubmatch(path)
					if len(match) > 0 {
						id := match[1]
						fmt.Println("ID:", id)
						ctx.Request.URL.Path = re.ReplaceAllString(path, nodeRule.TargetRoute)
						proxy.ServeHTTP(ctx.Writer, ctx.Request)
						return
					} else {
						fmt.Println("No match")
					}
				}

				http.StripPrefix(r, handler).ServeHTTP(ctx.Writer, ctx.Request)
			})
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
