package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/grape/internal/controller/v1"
)

type UserRouter struct {
	Name    string
	control *v1.UserController
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		Name:    "user",
		control: v1.NewUserController(),
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
