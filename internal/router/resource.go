package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type ResourceRouter struct {
	Name    string
	control *v1.ResourceController
}

func NewResourceRouter() *ResourceRouter {
	return &ResourceRouter{
		Name:    string(global.RGN_resource),
		control: &v1.ResourceController{},
	}
}

func (router ResourceRouter) Register(group *gin.RouterGroup) {
	f := group.Group(router.Name)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT("", router.control.Modify)
	f.DELETE(":id", router.control.Del)
	f.POST("upload", router.control.UploadFile)
	f.DELETE("upload/:name", router.control.UnUploadFile)
	f.GET("upload/sse", router.control.UploadFileSSE)
}
