package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type CertController struct{}

// Create doc
//
//	@tags			证书信息
//	@Summary		创建证书信息
//	@Description	创建证书信息
//	@Produce		json
//	@Param			data	body		repository.CreateCert	true	"创建证书信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/cert [post]
//	@Security		ApiKeyAuth
func (CertController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreateCert)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Cert{}.Create(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}

// Modify doc
//
//	@tags			证书信息
//	@Summary		修改证书信息
//	@Description	修改证书信息
//	@Produce		json
//	@Param			data	body		repository.ModifyCert	true	"修改证书信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/cert [put]
//	@Security		ApiKeyAuth
func (CertController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.ModifyCert)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Cert{}.Modify(req)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	c.Success()
}

// Query doc
//
//	@tags			证书信息
//	@Summary		查询证书信息
//	@Description	查询证书信息
//	@Produce		json
//	@Param			params	query		repository.QueryCert	true	"查询条件"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/cert [get]
//	@Security		ApiKeyAuth
func (CertController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryCert)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	list, err := logic.Cert{}.Get(req)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			证书信息
//	@Summary		通过编码查询证书信息
//	@Description	通过编码查询证书信息
//	@Produce		json
//	@Param			id	path		string			true	"证书信息编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/cert/{id} [get]
//	@Security		ApiKeyAuth
func (CertController) GetById(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.Cert{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessData(info)
}

// Del doc
//
//	@tags			证书信息
//	@Summary		删除证书信息
//	@Description	删除证书信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/cert [delete]
//	@Security		ApiKeyAuth
func (CertController) Del(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.Cert{}.Del(id)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	c.Success()
}
