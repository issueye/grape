package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type MenuController struct {
	controller.Controller
}

func NewMenuController() *MenuController {
	return new(MenuController)
}

// List doc
//
//	@tags			菜单管理
//	@Summary		获取菜单列表
//	@Description	获取菜单列表
//	@Produce		json
//	@Param			params	query		repository.QueryMenu			true	"查询条件"
//	@Success		200		{object}	controller.Full{[]model.Menu}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base					"错误返回内容"
//	@Router			/api/v1/menu [get]
//	@Security		ApiKeyAuth
func (MenuController) List(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.QueryMenu)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定请求内容失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	list, err := service.NewMenu().List(req)
	if err != nil {
		control.FailByMsgf("查询菜单信息列表失败 %s", err.Error())
		return
	}

	control.SuccessAutoData(req, list)
}

// TreeList doc
//
//	@tags			菜单管理
//	@Summary		获取菜单列表-树形结构
//	@Description	获取菜单列表-树形结构
//	@Produce		json
//	@Param			params	query		repository.QueryMenu						true	"查询条件"
//	@Success		200		{object}	controller.Full{[]repository.ResTreeMenus}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base								"错误返回内容"
//	@Router			/api/v1/menu/tree [get]
//	@Security		ApiKeyAuth
func (MenuController) TreeList(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.QueryMenu)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定请求内容失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	list, err := service.NewMenu().TreeList(req)
	if err != nil {
		control.FailByMsgf("查询菜单信息列表失败 %s", err.Error())
		return
	}

	control.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			菜单管理
//	@Summary		通过编码获取菜单
//	@Description	通过编码获取菜单
//	@Produce		json
//	@Param			id	path		string								true	"id"
//	@Success		200	{object}	controller.Full{data=model.Menu}	"code: 200 成功"
//	@Failure		500	{object}	controller.Base						"错误返回内容"
//	@Router			/api/v1/menu/{id} [get]
//	@Security		ApiKeyAuth
func (MenuController) GetById(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改菜单信息[id]不能为空")
		return
	}

	data, err := service.NewMenu().GetById(id)
	if err != nil {
		control.FailByMsgf("查询菜单信息列表失败 %s", err.Error())
		return
	}

	control.SuccessData(data)
}

// Create doc
//
//	@tags			菜单管理
//	@Summary		添加菜单信息
//	@Description	添加菜单信息
//	@Produce		json
//	@Param			data	body		repository.CreateMenu	true	"添加信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/menu [post]
//	@Security		ApiKeyAuth
func (MenuController) Create(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.CreateMenu)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	err = service.NewMenu().Create(req)
	if err != nil {
		control.FailByMsgf("添加菜单信息失败 %s", err.Error())
		return
	}
	control.Success()
}

// Modify doc
//
//	@tags			菜单管理
//	@Summary		修改菜单信息
//	@Description	修改菜单信息
//	@Produce		json
//	@Param			id		path		string					true	"id"
//	@Param			data	body		repository.ModifyMenu	true	"修改信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/menu/{id} [put]
//	@Security		ApiKeyAuth
func (MenuController) Modify(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.ModifyMenu)
	err := ctx.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改菜单信息[id]不能为空")
		return
	}

	err = logic.Menu{}.Modify(id, req)
	if err != nil {
		control.FailByMsgf("修改菜单信息失败 %s", err.Error())
		return
	}

	control.Success()
}

// ModifyStatus doc
//
//	@tags			菜单管理
//	@Summary		修改菜单状态
//	@Description	修改菜单状态
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/menu/state/{id} [put]
//	@Security		ApiKeyAuth
func (MenuController) ModifyState(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("参数[id]不能为空")
		return
	}

	err := logic.Menu{}.ModifyState(id)
	if err != nil {
		control.FailByMsgf("修改用户状态失败 %s", err.Error())
		return
	}

	control.Success()
}

// Delete doc
//
//	@tags			菜单管理
//	@Summary		删除菜单信息
//	@Description	删除菜单信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/menu/{id} [delete]
//	@Security		ApiKeyAuth
func (MenuController) Delete(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.Menu{}.Delete(id)
	if err != nil {
		control.FailByMsgf("删除菜单信息失败 %s", err.Error())
		return
	}

	control.Success()
}
