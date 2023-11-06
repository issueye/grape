package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
)

type PortRouter struct {
	Name    string
	control *v1.PortController
}

func NewPortRouter() *PortRouter {
	return &PortRouter{
		Name:    "port",
		control: &v1.PortController{},
	}
}

func (router PortRouter) Register(group *gin.RouterGroup) {
	f := group.Group(router.Name)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.PUT("reload/:id", router.control.Reload)
	f.POST("", router.control.Create)
	f.PUT("", router.control.Modify)
	f.PUT("state/:id", router.control.ModifyState)
	f.DELETE(":id", router.control.Del)
}
