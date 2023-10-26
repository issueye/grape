package router

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/controller"
)

type UserRouter struct {
	Name    string
	control *controller.UserController
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		Name:    "user",
		control: controller.NewUserController(),
	}
}

func (user UserRouter) Register(group *gin.RouterGroup) {
	f := group.Group(user.Name)
	f.GET("", user.control.List)
	f.GET(":id", user.control.GetById)
	f.POST("", user.control.Create)
	f.PUT(":id", user.control.Modify)
	f.PUT("state/:id", user.control.ModifyStatus)
	f.DELETE(":id", user.control.Delete)
}
