package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

// QueryHttpMessages doc
//
//	@tags		    http报文信息
//	@Summary		查询HTTP请求信息
//	@Description	查询HTTP请求信息
//	@Produce		json
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/query/httpMessages [get]
//	@Security		ApiKeyAuth
func QueryHttpMessages(ctx *gin.Context) {
	c := controller.New(ctx)

	lc := &logic.Logic{}
	list, err := lc.GetHttpMessages(&repository.HttpRequest{})
	if err != nil {
		c.FailByMsgf("查询失败: %v", err)
		return
	}

	c.SuccessData(list)
}
