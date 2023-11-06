package v1

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type PortController struct{}

// Create doc
//
//	@tags			端口信息
//	@Summary		创建端口信息
//	@Description	创建端口信息
//	@Produce		json
//	@Param			data	body		repository.CreatePort	true	"创建端口信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/port [post]
//	@Security		ApiKeyAuth
func (PortController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreatePort)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Port{}.Create(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}

// Modify doc
//
//	@tags			端口信息
//	@Summary		修改端口信息
//	@Description	修改端口信息
//	@Produce		json
//	@Param			data	body		repository.ModifyPort	true	"修改端口信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/port [put]
//	@Security		ApiKeyAuth
func (PortController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.ModifyPort)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Port{}.Modify(req)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	c.Success()
}

// ModifyState doc
//
//	@tags			端口信息
//	@Summary		修改端口使用状态
//	@Description	修改端口使用状态
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/port/state/{id} [put]
//	@Security		ApiKeyAuth
func (PortController) ModifyState(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	_, err := logic.Port{}.ModifyState(id)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	info, err := logic.Port{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询端口信息 %s", err.Error())
		return
	}

	// 处理状态
	fmt.Println("当前端口状态", info.State)
	ch := &global.Port{Id: info.ID, Port: info.Port}

	if info.State {
		ch.Action = global.AT_START
	} else {
		ch.Action = global.AT_STOP
	}

	global.PortChan <- ch

	c.Success()
}

// Reload doc
//
//	@tags			端口信息
//	@Summary		重启端口对应的服务
//	@Description	重启端口对应的服务
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/port/reload/{id} [put]
//	@Security		ApiKeyAuth
func (PortController) Reload(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.Port{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询端口信息 %s", err.Error())
		return
	}

	// 将状态修改为启动
	err = logic.Port{}.Start(id)
	if err != nil {
		c.FailByMsgf("启用端口[%d]失败 %s", info.Port, err.Error())
		return
	}

	ch := &global.Port{Id: info.ID, Port: info.Port, Action: global.AT_RELOAD}
	global.PortChan <- ch

	c.Success()
}

// Query doc
//
//	@tags			端口信息
//	@Summary		查询端口信息
//	@Description	查询端口信息
//	@Produce		json
//	@Param			params	query		repository.QueryPort	true	"查询条件"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/port [get]
//	@Security		ApiKeyAuth
func (PortController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryPort)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	list, err := logic.Port{}.Get(req)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			端口信息
//	@Summary		通过编码查询端口信息
//	@Description	通过编码查询端口信息
//	@Produce		json
//	@Param			id	path		string			true	"端口信息编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/port/{id} [get]
//	@Security		ApiKeyAuth
func (PortController) GetById(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.Port{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessData(info)
}

// Del doc
//
//	@tags			端口信息
//	@Summary		删除端口信息
//	@Description	删除端口信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/port [delete]
//	@Security		ApiKeyAuth
func (PortController) Del(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.Port{}.Del(id)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	c.Success()
}
