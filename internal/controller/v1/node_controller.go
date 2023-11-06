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

type NodeController struct{}

// Create doc
//
//	@tags			节点信息
//	@Summary		创建节点信息
//	@Description	创建节点信息
//	@Produce		json
//	@Param			data	body		repository.CreateNode	true	"创建节点信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/node [post]
//	@Security		ApiKeyAuth
func (NodeController) Create(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.CreateNode)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Node{}.Create(req)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	// 返回成功
	c.Success()
}

// Modify doc
//
//	@tags			节点信息
//	@Summary		修改节点信息
//	@Description	修改节点信息
//	@Produce		json
//	@Param			data	body		repository.ModifyNode	true	"修改节点信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/node [put]
//	@Security		ApiKeyAuth
func (NodeController) Modify(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.ModifyNode)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	err = logic.Node{}.Modify(req)
	if err != nil {
		c.FailByMsgf("更新信息失败 %s", err.Error())
		return
	}

	c.Success()
}

// Query doc
//
//	@tags			节点信息
//	@Summary		查询节点信息
//	@Description	查询节点信息
//	@Produce		json
//	@Param			params	query		repository.QueryNode	true	"查询条件"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/node [get]
//	@Security		ApiKeyAuth
func (NodeController) Query(ctx *gin.Context) {
	c := controller.New(ctx)

	// 绑定请求数据
	req := new(repository.QueryNode)
	err := c.Bind(req)
	if err != nil {
		c.FailBind(err)
		return
	}

	list, err := logic.Node{}.Get(req)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			节点信息
//	@Summary		通过编码查询节点信息
//	@Description	通过编码查询节点信息
//	@Produce		json
//	@Param			id	path		string			true	"节点信息编码"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/node/{id} [get]
//	@Security		ApiKeyAuth
func (NodeController) GetById(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	info, err := logic.Node{}.GetById(id)
	if err != nil {
		c.FailByMsgf("查询失败 %s", err.Error())
		return
	}

	c.SuccessData(info)
}

// Del doc
//
//	@tags			节点信息
//	@Summary		删除节点信息
//	@Description	删除节点信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/node [delete]
//	@Security		ApiKeyAuth
func (NodeController) Del(ctx *gin.Context) {
	c := controller.New(ctx)

	id := c.Param("id")
	if id == "" {
		c.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.Node{}.Del(id)
	if err != nil {
		c.FailByMsg(err.Error())
		return
	}

	c.Success()
}

// uploadFile doc
//
//	@tags			节点信息
//	@Summary		上传静态页面
//	@Description	上传静态页面
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/node [post]
//	@Security		ApiKeyAuth
func (NodeController) UploadFile(ctx *gin.Context) {
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

	nodeInfo, err := logic.Node{}.GetById(nodeId)
	if err != nil {
		c.FailByMsgf("获取节点信息失败 %s", err.Error())
		return
	}

	// 获取文件名，并创建新的文件存储
	filename := header.Filename
	path := filepath.Join("runtime", "static", "pages", nodeInfo.PortId, nodeInfo.Name, filename)
	name := utils.FileGetName(path)
	// 解压之后的文件夹
	savePath := filepath.Join("runtime", "static", "pages", nodeInfo.PortId, nodeInfo.Name, name)
	fmt.Println("savePath", savePath)

	// 创建文件夹
	exists, err := utils.PathExists(savePath)
	if err != nil {
		c.FailByMsgf("创建文件夹失败 %s", err.Error())
		return
	}
	if !exists {
		c.FailByMsgf("创建文件夹失败 %s", err.Error())
		return
	}

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
	err = logic.Node{}.ModifyByMap(nodeId, map[string]any{"page_path": fmt.Sprintf("/%s/%s/", nodeInfo.Name, name), "file_name": name})
	if err != nil {
		c.FailByMsgf("更新节点信息失败 %s", err.Error())
		return
	}

	c.SuccessByMsg("文件上传成功")
}
