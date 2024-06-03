package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type GzipFilterController struct{}

// Create doc
//
//	@tags			gzip过滤信息
//	@Summary		创建gzip过滤信息
//	@Description	创建gzip过滤信息
//	@Produce		json
//	@Param			data	body		repository.CreateGzipFilter	true	"创建gzip过滤信息"
//	@Success		200		{object}	controller.Base				"code: 200 成功"
//	@Failure		500		{object}	controller.Base				"错误返回内容"
//	@Router			/api/v1/gzipFilter [post]
//	@Security		ApiKeyAuth
func (GzipFilterController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreateGzipFilter)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.GzipFilter{}.Create(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}

// Modify doc
//
//	@tags			gzip过滤信息
//	@Summary		修改gzip过滤信息
//	@Description	修改gzip过滤信息
//	@Produce		json
//	@Param			data	body		repository.ModifyGzipFilter	true	"修改gzip过滤信息"
//	@Success		200		{object}	controller.Base				"code: 200 成功"
//	@Failure		500		{object}	controller.Base				"错误返回内容"
//	@Router			/api/v1/gzipFilter/{id} [put]
//	@Security		ApiKeyAuth
func (GzipFilterController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	// 绑定请求数据
	req := new(repository.ModifyGzipFilter)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.GzipFilter{}.Modify(id, req)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	c.Success()
}

// Query doc
//
//	@tags			gzip过滤信息
//	@Summary		查询gzip过滤信息
//	@Description	查询gzip过滤信息
//	@Produce		json
//	@Param			params	query		repository.QueryGzipFilter	true	"查询条件"
//	@Success		200		{object}	controller.Base				"code: 200 成功"
//	@Failure		500		{object}	controller.Base				"错误返回内容"
//	@Router			/api/v1/gzipFilter [get]
//	@Security		ApiKeyAuth
func (GzipFilterController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryGzipFilter)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	if req.PortId == "" {
		c.Success()
	}

	list, err := logic.GzipFilter{}.Get(req)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			gzip过滤信息
//	@Summary		通过编码查询gzip过滤信息
//	@Description	通过编码查询gzip过滤信息
//	@Produce		json
//	@Param			id	path		string			true	"gzip过滤信息编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/gzipFilter/{id} [get]
//	@Security		ApiKeyAuth
func (GzipFilterController) GetById(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.GzipFilter{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessData(info)
}

// Del doc
//
//	@tags			gzip过滤信息
//	@Summary		删除gzip过滤信息
//	@Description	删除gzip过滤信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/gzipFilter/{id} [delete]
//	@Security		ApiKeyAuth
func (GzipFilterController) Del(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.GzipFilter{}.Del(id)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	c.Success()
}
