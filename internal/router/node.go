package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
)

type NodeRouter struct {
	Name    string
	control *v1.NodeController
}

func NewNodeRouter() *NodeRouter {
	return &NodeRouter{
		Name:    "node",
		control: &v1.NodeController{},
	}
}

func (router NodeRouter) Register(group *gin.RouterGroup) {
	f := group.Group(router.Name)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT("", router.control.Modify)
	f.DELETE(":id", router.control.Del)
}
