package engine

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/global"
)

// 存放 http server 对象
var Servers = new(sync.Map)

func Start() {
	go func() {
		for {
			select {
			case p := <-global.PortChan:
				{
					switch p.Action {
					case global.AT_START:
						RunServer(p.Id, p.Port)
					case global.AT_STOP:
						StopServer(p.Id, p.Port)
					case global.AT_RELOAD:
						ReloadServer(p.Id, p.Port)
					}
				}
			}
		}
	}()
}

func ReloadServer(portId string, port int) {
	global.Log.Infof("[%d]端口号开始重启...", port)

	StopServer(portId, port)
	RunServer(portId, port)
}

func StopServer(portId string, port int) {
	global.Log.Infof("[%d]端口号停用服务...", port)

	value, ok := Servers.Load(portId)
	if ok {
		server := value.(*http.Server)
		server.Shutdown(context.Background())

		// 删除对象
		Servers.Delete(portId)
	}
}

func runServer(portId string, port int) {
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		c := controller.New(ctx)
		c.SuccessByMsgf("端口号[%d]返回消息", port)
	})

	proxy := ReverseProxyHttpHandler("http://127.0.0.1:10070")

	engine.POST("/login", func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	engine.GET("/page/vueRouter", func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	engine.Any("/granada/api/v1/*path", func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	LoadNode(portId, engine)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}

	// 存放到 map 中
	Servers.Store(portId, server)

	err := server.ListenAndServe()
	if err != nil {
		global.Log.Errorf("启动服务失败 %s", err.Error())
	}
}

func RunServer(portId string, port int) {
	global.Log.Infof("[%d]端口号启用服务...", port)

	go runServer(portId, port)
}
