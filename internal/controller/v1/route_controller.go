package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type RouteController struct{}

// Create doc
//
//	@tags			路由匹配信息
//	@Summary		创建路由匹配信息
//	@Description	创建路由匹配信息
//	@Produce		json
//	@Param			data	body		repository.CreateRoute	true	"创建路由匹配信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/route [post]
//	@Security		ApiKeyAuth
func (RouteController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreateRoute)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
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
//	@tags			路由匹配信息
//	@Summary		修改路由匹配信息
//	@Description	修改路由匹配信息
//	@Produce		json
//	@Param			data	body		repository.ModifyRoute	true	"修改路由匹配信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/route [put]
//	@Security		ApiKeyAuth
func (RouteController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.ModifyRoute)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
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
//	@tags			路由匹配信息
//	@Summary		查询路由匹配信息
//	@Description	查询路由匹配信息
//	@Produce		json
//	@Param			params	query		repository.QueryRoute	true	"查询条件"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/route [get]
//	@Security		ApiKeyAuth
func (RouteController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryRoute)
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
//	@tags			路由匹配信息
//	@Summary		通过编码查询路由匹配信息
//	@Description	通过编码查询路由匹配信息
//	@Produce		json
//	@Param			id	path		string			true	"路由匹配信息编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/route/{id} [get]
//	@Security		ApiKeyAuth
func (RouteController) GetById(ctx *gin.Context) {
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
//	@tags			路由匹配信息
//	@Summary		删除路由匹配信息
//	@Description	删除路由匹配信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/route [delete]
//	@Security		ApiKeyAuth
func (RouteController) Del(ctx *gin.Context) {
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
