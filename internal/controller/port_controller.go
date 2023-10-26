package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/repository"
)

type PortController struct{}

// 创建节点
func (PortController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreatePortInfo)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}
}
