package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type RuleController struct{}

// Create doc
//
//	@tags			匹配规则信息
//	@Summary		创建匹配规则信息
//	@Description	创建匹配规则信息
//	@Produce		json
//	@Param			data	body		repository.CreateRule	true	"创建匹配规则信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/rule [post]
//	@Security		ApiKeyAuth
func (RuleController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreateRule)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	// 如果匹配类型为 GIN匹配时 需要检查是否符合规范
	if req.MatchType == 1 {
		err = logic.Page{}.CheckData(req.PortId, req.Name)
		if err != nil {
			c.FailByMsg(err.Error())
			return
		}
	}

	err = logic.Route{}.Create(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}

// Modify doc
//
//	@tags			匹配规则信息
//	@Summary		修改匹配规则信息
//	@Description	修改匹配规则信息
//	@Produce		json
//	@Param			data	body		repository.ModifyRule	true	"修改匹配规则信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/rule [put]
//	@Security		ApiKeyAuth
func (RuleController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.ModifyRule)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	// 如果匹配类型为 GIN匹配时 需要检查是否符合规范
	if req.MatchType == 1 {
		err = logic.Page{}.CheckData(req.PortId, req.Name)
		if err != nil {
			c.FailByMsg(err.Error())
			return
		}
	}

	err = logic.Route{}.Modify(req)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	c.Success()
}

// Query doc
//
//	@tags			匹配规则信息
//	@Summary		查询匹配规则信息
//	@Description	查询匹配规则信息
//	@Produce		json
//	@Param			params	query		repository.QueryRule	true	"查询条件"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/rule [get]
//	@Security		ApiKeyAuth
func (RuleController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryRule)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	list, err := logic.Route{}.Get(req)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			匹配规则信息
//	@Summary		通过编码查询匹配规则信息
//	@Description	通过编码查询匹配规则信息
//	@Produce		json
//	@Param			id	path		string			true	"匹配规则信息编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/rule/{id} [get]
//	@Security		ApiKeyAuth
func (RuleController) GetById(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.Route{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessData(info)
}

// Del doc
//
//	@tags			匹配规则信息
//	@Summary		删除匹配规则信息
//	@Description	删除匹配规则信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/rule [delete]
//	@Security		ApiKeyAuth
func (RuleController) Del(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.Route{}.Del(id)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	c.Success()
}
