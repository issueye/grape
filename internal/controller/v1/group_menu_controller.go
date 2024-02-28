package v1

import (
	"errors"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type GroupMenuController struct {
	controller.Controller
}

func NewGroupMenuController() *GroupMenuController {
	return new(GroupMenuController)
}

// List doc
//
//	@tags			菜单组管理
//	@Summary		获取菜单组列表
//	@Description	获取菜单组列表
//	@Produce		json
//	@Param			params	query		repository.QueryGroupMenu			true	"查询条件"
//	@Success		200		{object}	controller.Full{[]model.GroupMenu}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base						"错误返回内容"
//	@Router			/api/v1/groupMenu [get]
//	@Security		ApiKeyAuth
func (GroupMenuController) List(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.QueryGroupMenu)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定请求内容失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	list, err := service.NewGroupMenu().List(req)
	if err != nil {
		control.FailByMsgf("查询菜单组信息列表失败 %s", err.Error())
		return
	}

	control.SuccessAutoData(req, list)
}

// getMenu doc
//
//	@tags			菜单组管理
//	@Summary		获取菜单树列表
//	@Description	获取菜单树列表
//	@Produce		json
//	@Param			groupId	path		string								true	"id"
//	@Param			params	query		repository.QueryGroupMenu			true	"查询条件"
//	@Success		200		{object}	controller.Full{[]model.GroupMenu}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base						"错误返回内容"
//	@Router			/api/v1/groupMenu/getMenu/{groupId} [get]
//	@Security		ApiKeyAuth
func (GroupMenuController) GetMenu(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("groupId")
	if id == "" {
		control.FailByMsg("[groupId]不能为空")
		return
	}

	req := new(repository.QueryGroupMenu)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定请求内容失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	req.GroupId = id
	list, err := service.NewGroupMenu().List(req)
	if err != nil {
		control.FailByMsgf("查询菜单组信息列表失败 %s", err.Error())
		return
	}

	resDatas := make([]*repository.ResGroupMenu, 0)

	findFirstLayer := func(id string) *repository.ResGroupMenu {
		for _, element := range resDatas {
			if element.MenuId == id {
				return element
			}
		}

		resData := new(repository.ResGroupMenu)
		resData.MenuId = id
		resData.Children = make([]*repository.ResGroupMenu, 0)

		resDatas = append(resDatas, resData)
		return resData
	}

	findSecondLayer := func(parentId, id string) *repository.ResGroupMenu {
		menu := findFirstLayer(parentId)

		for _, element := range menu.Children {
			if element.MenuId == id {
				return element
			}
		}

		resData := new(repository.ResGroupMenu)
		resData.MenuId = id
		resData.Children = make([]*repository.ResGroupMenu, 0)
		menu.Children = append(menu.Children, resData)

		return resData
	}

	for _, element := range list {
		if !req.GetAll {
			if element.State == 0 && element.Level == 1 {
				continue
			}
		}

		if element.Level == 0 {
			menu := findFirstLayer(element.MenuId)
			if menu.Route != "" {
				continue
			}

			menu.ID = element.ID
			menu.MenuId = element.MenuId
			menu.Name = element.Name
			menu.Title = element.Title
			menu.Route = element.Route
			menu.Icon = element.Icon
			menu.Auth = element.Auth
			menu.State = element.State
		}

		if element.Level == 1 {
			menu := findSecondLayer(element.ParentId, element.MenuId)
			if menu.Route != "" {
				continue
			}

			menu.ID = element.ID
			menu.MenuId = element.MenuId
			menu.Name = element.Name
			menu.Title = element.Title
			menu.Route = element.Route
			menu.Icon = element.Icon
			menu.Auth = element.Auth
			menu.State = element.State
		}
	}

	// 处理一级菜单，如果一级菜单没有子菜单的则移除
	resDatas = slices.DeleteFunc(resDatas, func(rgm *repository.ResGroupMenu) bool {
		return len(rgm.Children) == 0 && rgm.Title != "首页"
	})

	control.SuccessData(resDatas)
}

// GetById doc
//
//	@tags			菜单组管理
//	@Summary		通过编码获取菜单组
//	@Description	通过编码获取菜单组
//	@Produce		json
//	@Param			id	path		string									true	"id"
//	@Success		200	{object}	controller.Full{data=model.GroupMenu}	"code: 200 成功"
//	@Failure		500	{object}	controller.Base							"错误返回内容"
//	@Router			/api/v1/groupMenu/{id} [get]
//	@Security		ApiKeyAuth
func (GroupMenuController) GetById(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改菜单组信息[id]不能为空")
		return
	}

	data, err := service.NewGroupMenu().GetById(id)
	if err != nil {
		global.Log.Errorf("查询菜单组信息列表失败 %s", err.Error())
		control.FailByMsg("查询菜单组信息列表失败")
		return
	}

	control.SuccessData(data)
}

// Create doc
//
//	@tags			菜单组管理
//	@Summary		添加菜单组信息
//	@Description	添加菜单组信息
//	@Produce		json
//	@Param			data	body		repository.CreateGroupMenu	true	"添加信息"
//	@Success		200		{object}	controller.Base				"code: 200 成功"
//	@Failure		500		{object}	controller.Base				"错误返回内容"
//	@Router			/api/v1/groupMenu [post]
//	@Security		ApiKeyAuth
func (GroupMenuController) Create(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.CreateGroupMenu)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	err = service.NewGroupMenu().Create(req)
	if err != nil {
		control.FailByMsgf("添加菜单组信息失败 %s", err.Error())
		return
	}
	control.Success()
}

// Modify doc
//
//	@tags			菜单组管理
//	@Summary		修改菜单组信息
//	@Description	修改菜单组信息
//	@Produce		json
//	@Param			id		path		string						true	"id"
//	@Param			data	body		repository.ModifyGroupMenu	true	"修改信息"
//	@Success		200		{object}	controller.Base				"code: 200 成功"
//	@Failure		500		{object}	controller.Base				"错误返回内容"
//	@Router			/api/v1/groupMenu/{id} [put]
//	@Security		ApiKeyAuth
func (GroupMenuController) Modify(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.ModifyGroupMenu)
	err := ctx.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改菜单组信息[id]不能为空")
		return
	}

	err = logic.GroupMenu{}.Modify(id, req)
	if err != nil {
		control.FailByMsgf("修改菜单组信息失败 %s", err.Error())
		return
	}

	control.Success()
}

// ModifyGroupMenuState doc
//
//	@tags			菜单组管理
//	@Summary		修改菜单组状态信息
//	@Description	修改菜单组状态信息
//	@Produce		json
//	@Param			id		path		string							true	"id"
//	@Param			data	body		repository.ModifyGroupMenuState	true	"修改信息"
//	@Success		200		{object}	controller.Base					"code: 200 成功"
//	@Failure		500		{object}	controller.Base					"错误返回内容"
//	@Router			/api/v1/groupMenu/auth/{id} [put]
//	@Security		ApiKeyAuth
func (GroupMenuController) ModifyGroupMenuState(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.ModifyGroupMenuState)
	err := ctx.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改菜单组信息[id]不能为空")
		return
	}

	err = logic.GroupMenu{}.ModifyGroupMenuState(id, req)
	if err != nil {
		control.FailByMsgf("修改菜单组信息失败 %s", err.Error())
		return
	}

	control.Success()
}

// ModifyStatus doc
//
//	@tags			菜单组管理
//	@Summary		修改菜单组状态
//	@Description	修改菜单组状态
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/groupMenu/state/{id} [put]
//	@Security		ApiKeyAuth
func (GroupMenuController) ModifyState(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("参数[id]不能为空")
		return
	}

	err := logic.GroupMenu{}.ModifyState(id)
	if err != nil {
		control.FailByMsgf("修改用户状态失败 %s", err.Error())
		return
	}

	control.Success()
}

// Delete doc
//
//	@tags			菜单组管理
//	@Summary		删除菜单组信息
//	@Description	删除菜单组信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/groupMenu/{id} [delete]
//	@Security		ApiKeyAuth
func (GroupMenuController) Delete(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.GroupMenu{}.Delete(id)
	if err != nil {
		control.FailByMsgf("删除菜单组信息失败 %s", err.Error())
		return
	}

	control.Success()
}
