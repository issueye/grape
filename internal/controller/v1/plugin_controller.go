package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type PluginController struct {
	controller.Controller
}

func NewPluginController() *PluginController {
	return new(PluginController)
}

// Create doc
//
//	@tags			插件管理
//	@Summary		创建插件信息
//	@Description	创建插件信息
//	@Produce		json
//	@Param			data	body		repository.CreateRule	true	"插件信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/plugin [post]
//	@Security		ApiKeyAuth
func (PluginController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	req := new(repository.CreatePlugin)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = new(logic.Plugin).CreatePlugin(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}
