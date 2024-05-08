package v1

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/pkg/utils"
)

type ResourceController struct{}

// Create doc
//
//	@tags			页面信息
//	@Summary		创建页面信息
//	@Description	创建页面信息
//	@Produce		json
//	@Param			data	body		repository.CreateResource	true	"创建页面信息"
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
//	@tags			页面信息
//	@Summary		修改页面信息
//	@Description	修改页面信息
//	@Produce		json
//	@Param			data	body		repository.ModifyResource	true	"修改页面信息"
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
//	@tags			页面信息
//	@Summary		查询页面信息
//	@Description	查询页面信息
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
//	@tags			页面信息
//	@Summary		通过编码查询页面信息
//	@Description	通过编码查询页面信息
//	@Produce		json
//	@Param			id	path		string			true	"页面信息编码"
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
//	@tags			页面信息
//	@Summary		删除页面信息
//	@Description	删除页面信息
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
//	@tags			页面信息
//	@Summary		上传静态页面
//	@Description	上传静态页面
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/resource/upload [post]
//	@Security		ApiKeyAuth
func (ResourceController) UploadFile(ctx *gin.Context) {
	c := controller.New(ctx)
	//form表单
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.FailBind(err)
		return
	}

	nodeId := c.Request.FormValue("node_id")
	portId := c.Request.FormValue("port_id")
	if nodeId == "" || portId == "" {
		c.FailByMsg("[node_id]节点编码或者[port_id]端口编码不能为空")
		return
	}

	nodeInfo, err := logic.Resource{}.GetById(nodeId)
	if err != nil {
		c.FailByMsgf("获取页面信息失败 %s", err.Error())
		return
	}

	// 获取文件名，并创建新的文件存储
	filename := header.Filename
	path := filepath.Join("runtime", "static", "resources", nodeInfo.Folder, nodeInfo.Name, filename)
	name := utils.FileGetName(path)
	// 解压之后的文件夹
	savePath := filepath.Join("runtime", "static", "resources", nodeInfo.Folder, nodeInfo.Name, name)
	fmt.Println("savePath", savePath)

	// 创建文件夹
	utils.PathExists(savePath)

	// 创建上传文件
	out, err := os.Create(path)
	if err != nil {
		c.FailByMsgf("创建文件失败 %s", err.Error())
		return
	}
	defer out.Close()
	//将读取的文件流写到文件中
	_, err = io.Copy(out, file)
	if err != nil {
		c.FailByMsgf("读取失败 %s", err.Error())
		return
	}

	// 解压文件
	err = utils.Unzip(path, savePath)
	if err != nil {
		c.FailByMsgf("解压zip文件失败 %s", err.Error())
		return
	}

	// 更新节点状态
	err = logic.Resource{}.ModifyByMap(nodeId, map[string]any{"page_path": fmt.Sprintf("/%s/%s/", nodeInfo.Name, name), "file_name": name})
	if err != nil {
		c.FailByMsgf("更新页面信息失败 %s", err.Error())
		return
	}

	c.SuccessByMsg("文件上传成功")
}
