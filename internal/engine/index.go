package engine

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/middleware"
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

type muxHandler = func(http.ResponseWriter, *http.Request)

type GrapeEngine struct {
	PortId  string             // 端口号编码
	Port    int                // 端口号
	Engine  *gin.Engine        // gin
	Mux     *mux.Router        // mux 对象
	Rules   []*Rule            // 规则列表
	Customs []*CustomRouteRule // 节点规则列表
}

type Rule struct {
	Name    string                 `json:"name"`   // 匹配规则
	Target  string                 `json:"target"` // 目标地址
	Route   string                 `json:"route"`  // 路由
	Node    string                 `json:"node"`   // 节点
	Method  string                 `json:"method"` // 方法
	Proxy   *httputil.ReverseProxy `json:"proxy"`  // 代理转发
	Handler gin.HandlerFunc        // 方法
}

type CustomRouteRule struct {
	Name    string                 `json:"name"`   // 匹配规则
	Target  string                 `json:"target"` // 目标地址
	Route   string                 `json:"route"`  // 路由
	Node    string                 `json:"node"`   // 节点
	Method  string                 `json:"method"` // 方法
	Proxy   *httputil.ReverseProxy `json:"proxy"`  // 代理转发
	Handler muxHandler             // 方法
}

func NewGrapeEngine(portId string, port int) *GrapeEngine {
	en := &GrapeEngine{
		PortId:  portId,
		Port:    port,
		Engine:  gin.Default(),
		Rules:   make([]*Rule, 0),
		Customs: make([]*CustomRouteRule, 0),
	}

	en.Engine.Use(middleware.TrafficMiddleware(global.Log))
	en.Engine.Use(gin.Recovery())
	en.Engine.GET("/", func(ctx *gin.Context) {
		c := controller.New(ctx)
		c.SuccessByMsgf("端口号[%d]返回消息", port)
	})
	en.Mux = mux.NewRouter()
	return en
}

func (grape *GrapeEngine) Init() error {
	err := grape.GinRoutes()
	if err != nil {
		return err
	}

	err = grape.GinPages()
	if err != nil {
		return err
	}

	// 自定义路由
	return grape.CustomRoutes()
}

func (grape *GrapeEngine) GinPages() error {
	// 处理页面
	nodeList, err := logic.Node{}.Get(&repository.QueryNode{
		PortId: grape.PortId,
	})
	if err != nil {
		return err
	}

	for _, node := range nodeList {
		nodePage := grape.Engine.Group(node.Name)
		dir := filepath.Join("runtime", "static", "pages", node.PortId, node.Name, node.FileName)
		fmt.Println("静态文件路径：", dir)
		nodePage.Static("/web", dir)
	}

	return nil
}

func (grape *GrapeEngine) GinRoutes() error {
	// 获取匹配规则
	ruleList, err := service.NewRule().Query(&repository.QueryRule{
		PortId:    grape.PortId,
		NodeId:    "-",
		MatchType: 1,
	})

	if err != nil {
		return err
	}

	// 处理普通接口
	for _, rule := range ruleList {
		target, err := service.NewTarget().FindById(rule.TargetId)
		if err != nil {
			continue
		}

		route := rule.TargetRoute
		route = strings.ReplaceAll(route, "/:", "/")
		route = strings.ReplaceAll(route, "/*path", "/path")

		r := &Rule{
			Name:   rule.Name,
			Target: target.Name,
			Route:  route,
			Method: rule.Method,
		}

		r.Proxy = ReverseProxyHttpHandler(target.Name)
		r.Handler = func(ctx *gin.Context) {

			referer := ctx.Request.Header.Get("Referer")
			if referer != "" {
				if strings.HasSuffix(referer, "/lineAdmin/web/") {
					fmt.Println("referer", referer)
				}
			}

			if route != "" {
				route := ""
				for _, p := range ctx.Params {
					route = r.replace(p.Key, p.Value)
				}

				ctx.Request.URL.Path = route
			}

			r.Proxy.ServeHTTP(ctx.Writer, ctx.Request)
		}

		grape.Rules = append(grape.Rules, r)
	}

	// 注册 api
	for _, rule := range grape.Rules {
		switch strings.ToUpper(rule.Method) {
		case "POST":
			grape.Engine.POST(rule.Name, rule.Handler)
		case "GET":
			grape.Engine.GET(rule.Name, rule.Handler)
		case "PUT":
			grape.Engine.PUT(rule.Name, rule.Handler)
		case "PATCH":
			grape.Engine.PATCH(rule.Name, rule.Handler)
		case "DELETE":
			grape.Engine.DELETE(rule.Name, rule.Handler)
		case "ANY":
			grape.Engine.Any(rule.Name, rule.Handler)
		default:
			grape.Engine.Any(rule.Name, rule.Handler)
		}

	}

	return nil
}

func (rule *CustomRouteRule) replace(key, value string) string {
	v := value
	v = strings.TrimPrefix(v, "/")
	return strings.ReplaceAll(rule.Route, fmt.Sprintf("{%s}", key), v)
}

func (grape *GrapeEngine) CustomRoutes() error {
	ruleList, err := service.NewRule().Query(&repository.QueryRule{
		PortId:    grape.PortId,
		MatchType: 2,
	})
	if err != nil {
		return err
	}

	for _, rule := range ruleList {
		custom := &CustomRouteRule{
			Name:   rule.Name,
			Target: rule.Target,
			Route:  rule.TargetRoute,
			Method: rule.Method,
		}

		custom.Proxy = ReverseProxyHttpHandler(custom.Target)
		custom.Handler = func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			route := ""
			if len(vars) == 0 {
				route = custom.Route
			} else {
				for key, val := range vars {
					route = custom.replace(key, val)
				}
			}

			r.URL.Path = route
			custom.Proxy.ServeHTTP(w, r)
		}

		grape.Customs = append(grape.Customs, custom)
	}

	for _, custom := range grape.Customs {
		switch strings.ToUpper(custom.Method) {
		case "POST":
			grape.Mux.HandleFunc(custom.Name, custom.Handler).Methods("POST")
		case "GET":
			grape.Mux.HandleFunc(custom.Name, custom.Handler).Methods("GET")
		case "PUT":
			grape.Mux.HandleFunc(custom.Name, custom.Handler).Methods("PUT")
		case "PATCH":
			grape.Mux.HandleFunc(custom.Name, custom.Handler).Methods("PATCH")
		case "DELETE":
			grape.Mux.HandleFunc(custom.Name, custom.Handler).Methods("DELETE")
		case "ANY":
			grape.Mux.HandleFunc(custom.Name, custom.Handler)
		default:
			grape.Mux.HandleFunc(custom.Name, custom.Handler)
		}
	}

	return nil
}

func (grape *GrapeEngine) Run() error {

	// 处理未匹配上的接口
	grape.Engine.NoRoute(gin.WrapH(grape.Mux))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", grape.Port),
		Handler: grape.Engine,
	}

	// 存放到 map 中
	Servers.Store(grape.PortId, server)

	err := server.ListenAndServe()
	if err != nil {
		global.Log.Errorf("启动服务失败 %s", err.Error())
		return err
	}

	return nil
}

func (rule *Rule) replace(key, value string) string {
	v := value
	v = strings.TrimPrefix(v, "/")
	return strings.ReplaceAll(rule.Route, key, v)
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
	grape := NewGrapeEngine(portId, port)
	err := grape.Init()
	if err != nil {
		global.Log.Errorf("初始化失败 %s", err.Error())
		return
	}

	err = grape.Run()
	if err != nil {
		global.Log.Errorf("启动失败 %s", err.Error())
		return
	}
}

func RunServer(portId string, port int) {
	global.Log.Infof("[%d]端口号启用服务...", port)
	go runServer(portId, port)
}
