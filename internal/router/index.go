package router

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/config"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/middleware"
)

type IRouters interface {
	Register(group *gin.RouterGroup, auth gin.HandlerFunc)
}

func InitRouter(r *gin.Engine) {
	name := config.GetParam(config.CfgServerApiName, "api").String()
	v := config.GetParam(config.CfgServerApiVersion, "v1").String()

	apiName := r.Group(name)
	version := apiName.Group(v)
	global.Auth = middleware.NewAuth()

	// 用户鉴权 auth
	{
		version.POST("login", global.Auth.LoginHandler)
		version.GET("logout", global.Auth.LogoutHandler)
		version.GET("refreshToken", global.Auth.RefreshHandler)
	}

	// 鉴权
	// version.Use(global.Auth.MiddlewareFunc())
	registerVersionRouter(version, global.Auth.MiddlewareFunc(),
		NewUserRouter(),       // 用户
		NewUserGroupRouter(),  // 用户组
		NewMenuRouter(),       // 菜单
		NewMenuGroupRouter(),  // 用户组菜单
		NewPortRouter(),       // 端口号
		NewGzipFilterRouter(), // GZIP过滤
		NewPageRouter(),       // 页面
		NewRouteRouter(),      // 路由
		NewTargetRouter(),     // 服务地址
		NewCertRouter(),       // 证书
		NewResourceRouter(),   // 资源
	)

	// 查询 http 报文
	version.GET("query/httpMessages", v1.QueryHttpMessages)
}

// registerRouter 注册路由
func registerVersionRouter(r *gin.RouterGroup, auth gin.HandlerFunc, iRouters ...IRouters) {
	for _, iRouter := range iRouters {
		iRouter.Register(r, auth)
	}
}
