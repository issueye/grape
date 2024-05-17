package v1

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
)

type ResourceController struct{}

// Create doc
//
//	@tags			资源
//	@Summary		创建资源
//	@Description	创建资源
//	@Produce		json
//	@Param			data	body		repository.CreateResource	true	"创建资源"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/resource [post]
//	@Security		ApiKeyAuth
func (ResourceController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreateResource)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Resource{}.Create(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}

// Modify doc
//
//	@tags			资源
//	@Summary		修改资源
//	@Description	修改资源
//	@Produce		json
//	@Param			data	body		repository.ModifyResource	true	"修改资源"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/resource [put]
//	@Security		ApiKeyAuth
func (ResourceController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.ModifyResource)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Resource{}.Modify(req)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	c.Success()
}

// Query doc
//
//	@tags			资源
//	@Summary		查询资源
//	@Description	查询资源
//	@Produce		json
//	@Param			params	query		repository.QueryResource	true	"查询条件"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/resource [get]
//	@Security		ApiKeyAuth
func (ResourceController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryResource)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	list, err := logic.Resource{}.Get(req)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			资源
//	@Summary		通过编码查询资源
//	@Description	通过编码查询资源
//	@Produce		json
//	@Param			id	path		string			true	"资源编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/resource/{id} [get]
//	@Security		ApiKeyAuth
func (ResourceController) GetById(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.Resource{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessData(info)
}

// Del doc
//
//	@tags			资源
//	@Summary		删除资源
//	@Description	删除资源
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/resource [delete]
//	@Security		ApiKeyAuth
func (ResourceController) Del(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.Resource{}.Del(id)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	c.Success()
}

// uploadFile doc
//
//	@tags			资源
//	@Summary		上传资源
//	@Description	上传资源
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/resource/upload [post]
//	@Security		ApiKeyAuth
func (ResourceController) UploadFile(ctx *gin.Context) {
	c := controller.New(ctx)
	data := new(repository.UploadData)

	err := c.ShouldBind(data)
	if err != nil {
		c.FailByMsgf("绑定参数失败 %s", err.Error())
		return
	}

	resData, err := logic.Resource{}.Upload(data)
	if err != nil {
		c.FailByMsgf("上传文件失败")
		return
	}

	c.SuccessData(resData)
}

// UploadFileSSE doc
//
//	@tags			资源
//	@Summary		上传资源并且等待服务端返回
//	@Description	上传资源并且等待服务端返回
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/resource/upload/sse [get]
//	@Security		ApiKeyAuth
func (ResourceController) UploadFileSSE(ctx *gin.Context) {
	c := controller.New(ctx)

	// 使用 sse 代理此请求
	global.SSE.ServeHTTP(c.Writer, c.Request)
}

// uploadFile doc
//
//	@tags			资源
//	@Summary		移除资源
//	@Description	移除资源
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/resource/upload/{name} [delete]
//	@Security		ApiKeyAuth
func (ResourceController) UnUploadFile(ctx *gin.Context) {
	c := controller.New(ctx)

	name := c.Param("name")
	if name == "" {
		c.FailBind(errors.New("[name]不能为空"))
		return
	}

	ext := filepath.Ext(name)
	path := filepath.Join(global.GetResourcePathByType(ext), name)
	err := os.Remove(path)
	if err != nil {
		c.FailByMsgf("移除文件失败 %s", err.Error())
		return
	}

	c.SuccessByMsg("移除文件成功")
}
