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

type UserGroupController struct {
	controller.Controller
}

func NewUserGroupController() *UserGroupController {
	return new(UserGroupController)
}

// List doc
//
//	@tags			用户组信息管理
//	@Summary		获取用户组列表
//	@Description	获取用户组列表
//	@Produce		json
//	@Param			params	query		repository.QueryUserGroup				true	"查询条件"
//	@Success		200		{object}	controller.Full{[]model.UserGroupInfo}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base							"错误返回内容"
//	@Router			/api/v1/userGroup [get]
//	@Security		ApiKeyAuth
func (UserGroupController) List(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.QueryUserGroup)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定请求内容失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	list, err := service.NewUserGroup().List(req)
	if err != nil {
		global.Log.Errorf("查询用户信息列表失败 %s", err.Error())
		control.FailByMsg("查询用户信息列表失败")
		return
	}

	control.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			用户组信息管理
//	@Summary		通过编码获取用户组信息
//	@Description	通过编码获取用户组信息
//	@Produce		json
//	@Param			id	path		string										true	"id"
//	@Success		200	{object}	controller.Full{data=model.UserGroupInfo}	"code: 200 成功"
//	@Failure		500	{object}	controller.Base								"错误返回内容"
//	@Router			/api/v1/userGroup/{id} [get]
//	@Security		ApiKeyAuth
func (UserGroupController) GetById(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改用户信息[id]不能为空")
		return
	}

	data, err := service.NewUserGroup().GetById(id)
	if err != nil {
		global.Log.Errorf("查询用户信息列表失败 %s", err.Error())
		control.FailByMsg("查询用户信息列表失败")
		return
	}

	control.SuccessData(data)
}

// Create doc
//
//	@tags			用户组信息管理
//	@Summary		添加用户组信息
//	@Description	添加用户组信息
//	@Produce		json
//	@Param			data	body		repository.CreateUserGroup	true	"添加用户信息"
//	@Success		200		{object}	controller.Base				"code: 200 成功"
//	@Failure		500		{object}	controller.Base				"错误返回内容"
//	@Router			/api/v1/userGroup [post]
//	@Security		ApiKeyAuth
func (UserGroupController) Create(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.CreateUserGroup)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	err = service.NewUserGroup().Create(req)
	if err != nil {
		control.FailByMsgf("添加用户信息失败 %s", err.Error())
		return
	}
	control.Success()
}

// Modify doc
//
//	@tags			用户组信息管理
//	@Summary		修改用户组信息
//	@Description	修改用户组信息
//	@Produce		json
//	@Param			id		path		string						true	"id"
//	@Param			data	body		repository.ModifyUserGroup	true	"修改用户信息"
//	@Success		200		{object}	controller.Base				"code: 200 成功"
//	@Failure		500		{object}	controller.Base				"错误返回内容"
//	@Router			/api/v1/userGroup/{id} [put]
//	@Security		ApiKeyAuth
func (UserGroupController) Modify(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.ModifyUserGroup)
	err := ctx.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改用户信息[id]不能为空")
		return
	}

	err = service.NewUserGroup().Modify(req)
	if err != nil {
		control.FailByMsg("修改定时任务信息失败")
		return
	}

	control.Success()
}

// ModifyState doc
//
//	@tags			用户组信息管理
//	@Summary		修改用户组状态
//	@Description	修改用户组状态
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/userGroup/state/{id} [put]
//	@Security		ApiKeyAuth
func (UserGroupController) ModifyState(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("参数[id]不能为空")
		return
	}

	err := logic.UserGroup{}.ModifyState(id)
	if err != nil {
		control.FailByMsgf("修改用户组状态失败 %s", err.Error())
		return
	}

	control.Success()
}

// Delete doc
//
//	@tags			用户组信息管理
//	@Summary		删除用户组
//	@Description	删除用户组
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/userGroup/{id} [delete]
//	@Security		ApiKeyAuth
func (UserGroupController) Delete(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.UserGroup{}.Delete(id)
	if err != nil {
		control.FailByMsgf("删除用户组信息失败 %s", err.Error())
		return
	}

	control.Success()
}
