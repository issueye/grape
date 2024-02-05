package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type MenuRouter struct {
	Name    string
	control *v1.MenuController
}

func NewMenuRouter() *MenuRouter {
	return &MenuRouter{
		Name:    string(global.RGN_menu),
		control: v1.NewMenuController(),
	}
}

func (Menu MenuRouter) Register(group *gin.RouterGroup) {
	f := group.Group(Menu.Name)
	f.GET("", Menu.control.List)
	f.GET(":id", Menu.control.GetById)
	f.POST("", Menu.control.Create)
	f.PUT(":id", Menu.control.Modify)
	f.PUT("state/:id", Menu.control.ModifyState)
	f.DELETE(":id", Menu.control.Delete)
}

type MenuGroupRouter struct {
	Name    string
	control *v1.GroupMenuController
}

func NewMenuGroupRouter() *MenuGroupRouter {
	return &MenuGroupRouter{
		Name:    string(global.RGN_menuGroup),
		control: v1.NewGroupMenuController(),
	}
}

func (Menu *MenuGroupRouter) Register(group *gin.RouterGroup) {
	f := group.Group(Menu.Name)
	f.GET("", Menu.control.List)
	f.GET(":id", Menu.control.GetById)
	f.POST("", Menu.control.Create)
	f.PUT(":id", Menu.control.Modify)
	f.PUT("state/:id", Menu.control.ModifyState)
	f.DELETE(":id", Menu.control.Delete)
}
