package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type TargetController struct{}

// Create doc
//
//	@tags			目标地址信息
//	@Summary		创建目标地址信息
//	@Description	创建目标地址信息
//	@Produce		json
//	@Param			data	body		repository.CreateTarget	true	"创建目标地址信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/target [post]
//	@Security		ApiKeyAuth
func (TargetController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreateTarget)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Target{}.Create(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}

// Modify doc
//
//	@tags			目标地址信息
//	@Summary		修改目标地址信息
//	@Description	修改目标地址信息
//	@Produce		json
//	@Param			data	body		repository.ModifyTarget	true	"修改目标地址信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/target [put]
//	@Security		ApiKeyAuth
func (TargetController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.ModifyTarget)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Target{}.Modify(req)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	c.Success()
}

// Query doc
//
//	@tags			目标地址信息
//	@Summary		查询目标地址信息
//	@Description	查询目标地址信息
//	@Produce		json
//	@Param			params	query		repository.QueryTarget	true	"查询条件"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/target [get]
//	@Security		ApiKeyAuth
func (TargetController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryTarget)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	list, err := logic.Target{}.Get(req)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			目标地址信息
//	@Summary		通过编码查询目标地址信息
//	@Description	通过编码查询目标地址信息
//	@Produce		json
//	@Param			id	path		string			true	"目标地址信息编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/target/{id} [get]
//	@Security		ApiKeyAuth
func (TargetController) GetById(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.Target{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessData(info)
}

// Del doc
//
//	@tags			目标地址信息
//	@Summary		删除目标地址信息
//	@Description	删除目标地址信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/target [delete]
//	@Security		ApiKeyAuth
func (TargetController) Del(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.Target{}.Del(id)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	c.Success()
}
