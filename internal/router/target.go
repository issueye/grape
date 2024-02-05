package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type TargetRouter struct {
	Name    string
	control *v1.TargetController
}

func NewTargetRouter() *TargetRouter {
	return &TargetRouter{
		Name:    string(global.RGN_target),
		control: &v1.TargetController{},
	}
}

func (router TargetRouter) Register(group *gin.RouterGroup) {
	f := group.Group(router.Name)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT("", router.control.Modify)
	f.DELETE(":id", router.control.Del)
}
