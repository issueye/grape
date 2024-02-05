package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
	"github.com/issueye/grape/internal/global"
)

type UserRouter struct {
	Name    string
	control *v1.UserController
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		Name:    string(global.RGN_user),
		control: v1.NewUserController(),
	}
}

func (user UserRouter) Register(group *gin.RouterGroup) {
	f := group.Group(user.Name)
	f.GET("", user.control.List)
	f.GET(":id", user.control.GetById)
	f.POST("", user.control.Create)
	f.PUT(":id", user.control.Modify)
	f.PUT("state/:id", user.control.ModifyState)
	f.DELETE(":id", user.control.Delete)
}

type UserGroupRouter struct {
	Name    string
	control *v1.UserGroupController
}

func NewUserGroupRouter() *UserGroupRouter {
	return &UserGroupRouter{
		Name:    string(global.RGN_userGroup),
		control: v1.NewUserGroupController(),
	}
}

func (user *UserGroupRouter) Register(group *gin.RouterGroup) {
	f := group.Group(user.Name)
	f.GET("", user.control.List)
	f.GET(":id", user.control.GetById)
	f.POST("", user.control.Create)
	f.PUT(":id", user.control.Modify)
	f.PUT("state/:id", user.control.ModifyState)
	f.DELETE(":id", user.control.Delete)
}
