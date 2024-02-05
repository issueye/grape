package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type RouteRouter struct {
	Name    string
	control *v1.RuleController
}

func NewRouteRouter() *RouteRouter {
	return &RouteRouter{
		Name:    string(global.RGN_rule),
		control: &v1.RuleController{},
	}
}

func (router RouteRouter) Register(group *gin.RouterGroup) {
	f := group.Group(router.Name)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT("", router.control.Modify)
	f.DELETE(":id", router.control.Del)
}
