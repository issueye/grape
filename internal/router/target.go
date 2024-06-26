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

func (router TargetRouter) Register(group *gin.RouterGroup, auth gin.HandlerFunc) {
	f := group.Group(router.Name, auth)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT(":id", router.control.Modify)
	f.PUT("state/:id", router.control.ModifyState)
	f.DELETE(":id", router.control.Del)
}
