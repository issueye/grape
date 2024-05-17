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

func (router ResourceRouter) Register(group *gin.RouterGroup, auth gin.HandlerFunc) {
	f := group.Group(router.Name)
	f.GET("", auth, router.control.Query)
	f.GET(":id", auth, router.control.GetById)
	f.POST("", auth, router.control.Create)
	f.PUT("", auth, router.control.Modify)
	f.DELETE(":id", auth, router.control.Del)
	f.POST("upload", auth, router.control.UploadFile)
	f.DELETE("upload/:name", auth, router.control.UnUploadFile)
	f.GET("upload/sse", router.control.UploadFileSSE)
}
