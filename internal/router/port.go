package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type PortRouter struct {
	Name    string
	control *v1.PortController
}

func NewPortRouter() *PortRouter {
	return &PortRouter{
		Name:    string(global.RGN_port),
		control: &v1.PortController{},
	}
}

func (router PortRouter) Register(group *gin.RouterGroup, auth gin.HandlerFunc) {
	f := group.Group(router.Name, auth)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT(":id", router.control.Modify)
	f.PUT("start/:id", router.control.Start)
	f.PUT("stop/:id", router.control.Stop)
	f.PUT("state/:id", router.control.ModifyState)
	f.PUT("reload/:id", router.control.Reload)
	f.DELETE(":id", router.control.Del)
}
