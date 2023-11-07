package engine

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
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
	mr := mux.NewRouter()
	engine := gin.Default()
	// engine.GET("/", func(ctx *gin.Context) {
	// 	c := controller.New(ctx)
	// 	c.SuccessByMsgf("端口号[%d]返回消息", port)
	// })

	ruleList, err := service.NewRule().Query(&repository.QueryRule{
		PortId: portId,
		NodeId: "-",
	})

	if err != nil {
		return
	}

	for _, rule := range ruleList {
		target, err := service.NewTarget().FindById(rule.TargetId)
		if err != nil {
			continue
		}
		route := rule.TargetRoute

		proxy := ReverseProxyHttpHandler(target.Name)
		mr.HandleFunc(rule.Name, func(w http.ResponseWriter, r *http.Request) {
			if route != "" {
				r.URL.Path = route
			}
			proxy.ServeHTTP(w, r)
		})
		continue

		// engine.Any(rule.Name, func(ctx *gin.Context) {
		// 	if route != "" {
		// 		flag := "/*path"
		// 		fmt.Println("route[len(route)-len(flag):]", route[len(route)-len(flag):])
		// 		if route[len(route)-len(flag):] == flag {
		// 			path := ctx.Param("path")
		// 			ctx.Request.URL.Path = strings.ReplaceAll(route, flag, path)
		// 		} else {
		// 			ctx.Request.URL.Path = route
		// 		}
		// 	}

		// 	fmt.Println("ctx.Request.URL.Path", ctx.Request.URL.Path)
		// 	proxy.ServeHTTP(ctx.Writer, ctx.Request)
		// })
	}

	LoadNode(portId, engine)

	// mr.Handle("/", engine)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mr,
	}

	// 存放到 map 中
	Servers.Store(portId, server)

	err = server.ListenAndServe()
	if err != nil {
		global.Log.Errorf("启动服务失败 %s", err.Error())
	}
}

func RunServer(portId string, port int) {
	global.Log.Infof("[%d]端口号启用服务...", port)

	go runServer(portId, port)
}
