package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type GzipFilterRouter struct {
	Name    string
	control *v1.GzipFilterController
}

func NewGzipFilterRouter() *GzipFilterRouter {
	return &GzipFilterRouter{
		Name:    string(global.RGN_gzipFilter),
		control: &v1.GzipFilterController{},
	}
}

func (router GzipFilterRouter) Register(group *gin.RouterGroup, auth gin.HandlerFunc) {
	f := group.Group(router.Name, auth)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT(":id", router.control.Modify)
	f.DELETE(":id", router.control.Del)
}
