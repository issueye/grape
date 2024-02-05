package router

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/config"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/middleware"
)

type IRouters interface {
	Register(group *gin.RouterGroup)
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
	version.Use(global.Auth.MiddlewareFunc())
	registerVersionRouter(version,
		NewUserRouter(),      // 用户
		NewUserGroupRouter(), // 用户组
		NewMenuRouter(),      // 菜单
		NewMenuGroupRouter(), // 用户组菜单
		NewPortRouter(),      // 端口号
		NewNodeRouter(),      // 节点
		NewRouteRouter(),     // 路由
		NewTargetRouter(),    // 服务地址
		NewCertRouter(),      // 证书
	)
}

// registerRouter 注册路由
func registerVersionRouter(r *gin.RouterGroup, iRouters ...IRouters) {
	for _, iRouter := range iRouters {
		iRouter.Register(r)
	}
}
