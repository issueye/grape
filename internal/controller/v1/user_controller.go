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

type UserController struct {
	controller.Controller
}

func NewUserController() *UserController {
	return new(UserController)
}

// List doc
//
//	@tags			用户信息管理
//	@Summary		获取定时任务列表
//	@Description	获取定时任务列表
//	@Produce		json
//	@Param			params	query		repository.QueryUser				true	"查询条件"
//	@Success		200		{object}	controller.Full{[]model.UserInfo}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base						"错误返回内容"
//	@Router			/api/v1/user [get]
//	@Security		ApiKeyAuth
func (UserController) List(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.QueryUser)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定请求内容失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	list, err := service.NewUser().List(req)
	if err != nil {
		global.Log.Errorf("查询用户信息列表失败 %s", err.Error())
		control.FailByMsg("查询用户信息列表失败")
		return
	}

	control.SuccessAutoData(req, list)
}

// GetById doc
//
//	@tags			用户信息管理
//	@Summary		获取定时任务列表
//	@Description	获取定时任务列表
//	@Produce		json
//	@Param			id	path		string									true	"id"
//	@Success		200	{object}	controller.Full{data=model.UserInfo}	"code: 200 成功"
//	@Failure		500	{object}	controller.Base							"错误返回内容"
//	@Router			/api/v1/user/{id} [get]
//	@Security		ApiKeyAuth
func (UserController) GetById(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改用户信息[id]不能为空")
		return
	}

	data, err := service.NewUser().GetById(id)
	if err != nil {
		global.Log.Errorf("查询用户信息列表失败 %s", err.Error())
		control.FailByMsg("查询用户信息列表失败")
		return
	}

	control.SuccessData(data)
}

// Create doc
//
//	@tags			用户信息管理
//	@Summary		添加用户信息
//	@Description	添加用户信息
//	@Produce		json
//	@Param			data	body		repository.CreateUser	true	"添加用户信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/user [post]
//	@Security		ApiKeyAuth
func (UserController) Create(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.CreateUser)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	err = service.NewUser().Create(req)
	if err != nil {
		control.FailByMsgf("添加用户信息失败 %s", err.Error())
		return
	}
	control.Success()
}

// Modify doc
//
//	@tags			用户信息管理
//	@Summary		修改用户信息
//	@Description	修改用户信息
//	@Produce		json
//	@Param			id		path		string					true	"id"
//	@Param			data	body		repository.ModifyUser	true	"修改用户信息"
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/user/{id} [put]
//	@Security		ApiKeyAuth
func (UserController) Modify(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(repository.ModifyUser)
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

	err = logic.User{}.Modify(id, req)
	if err != nil {
		control.FailByMsgf("修改用户信息失败 %s", err.Error())
		return
	}

	control.Success()
}

// ModifyStatus doc
//
//	@tags			用户信息管理
//	@Summary		修改用户状态
//	@Description	修改用户状态
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/user/state/{id} [put]
//	@Security		ApiKeyAuth
func (UserController) ModifyState(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("参数[id]不能为空")
		return
	}

	err := logic.User{}.ModifyState(id)
	if err != nil {
		control.FailByMsgf("修改用户状态失败 %s", err.Error())
		return
	}

	control.Success()
}

// Delete doc
//
//	@tags			用户信息管理
//	@Summary		删除用户信息
//	@Description	删除用户信息
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Success		200	{object}	controller.Base	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/user/{id} [delete]
//	@Security		ApiKeyAuth
func (UserController) Delete(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailBind(errors.New("[id]不能为空"))
		return
	}

	err := logic.User{}.Delete(id)
	if err != nil {
		control.FailByMsgf("删除用户信息失败 %s", err.Error())
		return
	}

	control.Success()
}
