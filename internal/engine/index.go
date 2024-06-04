package engine

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/common/model"
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
						RunServer(p.PortInfo)
					case global.AT_STOP:
						StopServer(p.PortInfo)
					case global.AT_RELOAD:
						ReloadServer(p.PortInfo)
					}
				}
			}
		}
	}()
}

type muxHandler = func(http.ResponseWriter, *http.Request)

type GrapeEngine struct {
	PortId  string             // 端口号编码
	UseGzip bool               // 使用GZIP
	Port    int                // 端口号
	Engine  *gin.Engine        // gin
	Mux     *mux.Router        // mux 对象
	Rules   []*Rule            // 规则列表
	Customs []*CustomRouteRule // 节点规则列表
}

// type Transmit struct {
// 	TargetUrl string                 `json:"url"`     // 地址
// 	Count     uint64                 `json:"count"`   // 转发次数
// 	InFlow    uint64                 `json:"inFlow"`  // 入流量
// 	OutFlow   uint64                 `json:"outFlow"` // 出流量
// 	Proxy     *httputil.ReverseProxy `json:"proxy"`   // 代理转发
// 	lock      *sync.Mutex
// }

// func (tran *Transmit) calculateHTTPTraffic(req *http.Request, resp *http.Response) (int64, int64) {
// 	tran.lock.Lock()
// 	defer tran.lock.Unlock()

// 	var inBytes, outBytes int64
// 	inBytes += int64(len(req.Method)) + int64(len(req.URL.String()))
// 	for k, v := range req.Header {
// 		inBytes += int64(len(k)) + int64(len(v[0]))
// 	}
// 	inBytes += int64(req.ContentLength)

// 	for k, v := range resp.Header {
// 		outBytes += int64(len(k)) + int64(len(v[0]))
// 	}
// 	outBytes += int64(resp.ContentLength)

// 	func() {
// 		tran.lock.Lock()
// 		defer tran.lock.Unlock()

// 		tran.InFlow += uint64(inBytes)
// 		tran.OutFlow += uint64(outBytes)
// 	}()

// 	return inBytes, outBytes
// }

type Rule struct {
	Name    string                 `json:"name"`   // 匹配规则
	Target  string                 `json:"target"` // 目标地址
	Route   string                 `json:"route"`  // 路由
	Page    string                 `json:"Page"`   // 节点
	Method  string                 `json:"method"` // 方法
	Proxy   *httputil.ReverseProxy `json:"proxy"`  // 代理转发
	Handler gin.HandlerFunc        // 方法
}

type CustomRouteRule struct {
	Name    string                 `json:"name"`   // 匹配规则
	Target  string                 `json:"target"` // 目标地址
	Route   string                 `json:"route"`  // 路由
	Page    string                 `json:"Page"`   // 节点
	Method  string                 `json:"method"` // 方法
	Proxy   *httputil.ReverseProxy `json:"proxy"`  // 代理转发
	Handler muxHandler             // 方法
}

func NewGrapeEngine(port model.PortInfo) *GrapeEngine {
	en := &GrapeEngine{
		PortId:  port.ID,
		Port:    port.Port,
		UseGzip: port.UseGzip,
		Engine:  gin.Default(),
		Rules:   make([]*Rule, 0),
		Customs: make([]*CustomRouteRule, 0),
	}

	en.Engine.Use(middleware.TrafficMiddleware(global.Log))
	// en.Engine.Use(gzip.Gzip(gzip.DefaultCompression))
	en.Engine.Use(gin.Recovery())
	en.Engine.GET("/", func(ctx *gin.Context) {
		c := controller.New(ctx)
		c.SuccessByMsgf("端口号[%d]返回消息", port.Port)
	})
	en.Mux = mux.NewRouter()
	return en
}

func (grape *GrapeEngine) Init() error {

	if grape.UseGzip {
		err := grape.GinGzipFilter()
		if err != nil {
			return err
		}
	}

	err := grape.GinPages()
	if err != nil {
		return err
	}

	err = grape.GinRoutes()
	if err != nil {
		return err
	}

	// 自定义路由
	return grape.CustomRoutes()
}

type Page struct {
	RoutePath  string // 路径
	StaticPath string // 静态资源路径
}

func (grape *GrapeEngine) GinGzipFilter() error {
	list, err := logic.GzipFilter{}.Get(&repository.QueryGzipFilter{
		PortId: grape.PortId,
	})
	if err != nil {
		return err
	}

	extensions := []string{}
	paths := []string{}
	regexs := []string{}
	for _, filter := range list {
		switch filter.MatchType {
		case 1:
			paths = append(paths, filter.MatchContent)
		case 2:
			extensions = append(extensions, filter.MatchContent)
		case 3:
			regexs = append(regexs, filter.MatchContent)
		}
	}

	options := make([]gzip.Option, 0)

	if len(extensions) > 0 {
		options = append(options, gzip.WithExcludedExtensions(extensions))
	}

	if len(paths) > 0 {
		options = append(options, gzip.WithExcludedPaths(paths))
	}

	if len(regexs) > 0 {
		options = append(options, gzip.WithExcludedPathsRegexs(regexs))
	}

	if len(options) > 0 {
		grape.Engine.Use(gzip.Gzip(gzip.DefaultCompression, options...))
	} else {
		grape.Engine.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	return nil
}

// GinPages
// 加载页面
func (grape *GrapeEngine) GinPages() error {
	// 处理页面
	pageList, err := logic.Page{}.Get(&repository.QueryPage{
		PortId: grape.PortId,
	})
	if err != nil {
		return err
	}

	for _, page := range pageList {
		versionInfo, err := service.NewPage().FindByVersion(page.PortId, page.ProductCode, page.Version)
		if err != nil {
			global.Log.Errorf("页面[%s]未找到激活版本[%s] %s", page.Title, page.Version, err.Error())
			continue
		}
		// 在使用版本路由
		path := ""
		if page.UseVersionRoute == 1 {
			path = fmt.Sprintf("/%s/%s", page.Name, page.Version)
		} else {
			path = fmt.Sprintf("/%s", page.Name)
		}

		grape.Engine.Static(path, versionInfo.PagePath)

		// if page.UseGzip == 1 {
		// 	p.Use(gzip.Gzip(gzip.DefaultCompression))
		// }

		// grape.Engine.Static(path, versionInfo.PagePath).Use(gzip.Gzip(gzip.DefaultCompression))
	}

	return nil
}

func (grape *GrapeEngine) GinRoutes() error {
	// 获取匹配规则
	ruleList, err := service.NewRule().Query(&repository.QueryRule{
		PortId:    grape.PortId,
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
			if route != "" {
				// route := ""
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
	grape.Engine.NoRoute(func(ctx *gin.Context) {
		fmt.Println("GIN未匹配上的路由", "ctx.Request.URL.Path", ctx.Request.URL.Path)
		ctx.Next()
	}, gin.WrapH(grape.Mux))

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

func ReloadServer(port model.PortInfo) {
	global.Log.Infof("[%d]端口号开始重启...", port.Port)

	StopServer(port)
	RunServer(port)
}

func StopServer(port model.PortInfo) {
	global.Log.Infof("[%d]端口号停用服务...", port)

	value, ok := Servers.Load(port.ID)
	if ok {
		server := value.(*http.Server)
		server.Shutdown(context.Background())

		// 删除对象
		Servers.Delete(port.ID)
	}
}

func runServer(port model.PortInfo) {
	grape := NewGrapeEngine(port)
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

func RunServer(port model.PortInfo) {
	global.Log.Infof("[%d]端口号启用服务...", port.Port)
	go runServer(port)
}
