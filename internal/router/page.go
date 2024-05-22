package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type PageRouter struct {
	Name    string
	control *v1.PageController
}

func NewPageRouter() *PageRouter {
	return &PageRouter{
		Name:    string(global.RGN_page),
		control: &v1.PageController{},
	}
}

func (router PageRouter) Register(group *gin.RouterGroup, auth gin.HandlerFunc) {
	f := group.Group(router.Name, auth)
	f.GET("", router.control.Query)
	f.GET(":id", router.control.GetById)
	f.POST("", router.control.Create)
	f.PUT("", router.control.Modify)
	f.DELETE(":id", router.control.Del)

	f.GET("version/:productCode", router.control.GetPageVersinList)
}
