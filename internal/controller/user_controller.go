package controller

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/issueye/grape/internal/common/controller"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/model"
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
//	@Param			isNotPaging	query		string								false	"是否需要分页， 默认需要， 如果不分页 传 true"
//	@Param			pageNum		query		string								true	"页码， 如果不分页 传 0"
//	@Param			pageSize	query		string								true	"一页大小， 如果不分页 传 0"
//	@Param			crwmc		query		string								false	"任务名称"
//	@Param			desc		query		string								false	"任务描述"
//	@Success		200			{object}	controller.Full{[]models.TBZDDSRR}	true	"code: 200 成功"
//	@Failure		500			{object}	controller.Base						true	"错误返回内容"
//	@Router			/api/v1/users [get]
//	@Security		ApiKeyAuth
func (UserController) List(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(model.QueryUser)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定请求内容失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	list, err := service.NewUser(global.DB).List(req)
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
//	@Success		200	{object}	controller.Full{data=[]models.TBZDDSRR}	true	"code: 200 成功"
//	@Failure		500	{object}	controller.Base							true	"错误返回内容"
//	@Router			/api/v1/user/{id} [get]
//	@Security		ApiKeyAuth
func (UserController) GetById(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改用户信息[id]不能为空")
		return
	}

	idCondition, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		control.FailByMsg("[id]数据类型错误")
		return
	}

	data, err := service.NewUser(global.DB).GetById(idCondition)
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
//	@Param			data	body		model.CreateTBZDDSRR	true	"添加用户信息"
//	@Success		200		{object}	controller.Base			true	"code: 200 成功"
//	@Failure		500		{object}	controller.Base			true	"错误返回内容"
//	@Router			/api/v1/user [post]
//	@Security		ApiKeyAuth
func (UserController) Create(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(model.CreateUser)
	err := control.Bind(req)
	if err != nil {
		global.Log.Errorf("绑定参数失败 %s", err.Error())
		control.FailBind(err)
		return
	}

	err = service.NewUser(global.DB).Create(req)
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
//	@Param			data	body		model.ModifyTBZDDSRR	true	"修改用户信息"
//	@Success		200		{object}	controller.Base			true	"code: 200 成功"
//	@Failure		500		{object}	controller.Base			true	"错误返回内容"
//	@Router			/api/v1/user/{id} [put]
//	@Security		ApiKeyAuth
func (UserController) Modify(ctx *gin.Context) {
	control := controller.New(ctx)

	req := new(model.ModifyUser)
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

	idCondition, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		control.FailByMsg("[id]数据类型错误")
		return
	}

	// 查询定时任务
	info, err := service.NewUser(global.DB).GetById(idCondition)
	if err != nil {
		control.FailByMsg("查询定时任务失败")
		return
	}

	// 系统任务不允许修改表述信息
	if info.Mark == global.SYS_AUTO_CREATE {
		// 判断描述是否被修改
		if !strings.EqualFold(info.Mark, req.Mark) {
			control.FailByMsgf("【%s-%d】由系统生成, 不允许修改描述信息", info.Name, info.ID)
			return
		}
	}

	err = service.NewUser(global.DB).Modify(req)
	if err != nil {
		control.FailByMsg("修改定时任务信息失败")
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
//	@Success		200	{object}	controller.Base	true	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	true	"错误返回内容"
//	@Router			/api/v1/user/state/{id} [put]
//	@Security		ApiKeyAuth
func (UserController) ModifyStatus(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailByMsg("修改定时任务状态, 参数[id]不能为空")
		return
	}

	idCondition, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		control.FailByMsg("[id]数据类型错误")
		return
	}

	// 获取当前定时任务的状态
	info, err := service.NewUser(global.DB).GetById(idCondition)
	if err != nil {
		control.FailByMsgf("查询用户信息失败 %s", err.Error())
		return
	}

	err = service.NewUser(global.DB).Status(&model.StatusUser{
		ID:    idCondition,
		State: info.State,
	})
	if err != nil {
		control.FailByMsgf("修改用户信息失败 %s", err.Error())
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
//	@Success		200	{object}	controller.Base	true	"code: 200 成功"
//	@Failure		500	{object}	controller.Base	true	"错误返回内容"
//	@Router			/api/v1/user/{id} [delete]
//	@Security		ApiKeyAuth
func (UserController) Delete(ctx *gin.Context) {
	control := controller.New(ctx)

	id := control.Param("id")
	if id == "" {
		control.FailBind(errors.New("[id]不能为空"))
		return
	}

	idCondition, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		control.FailByMsg("[id]数据类型错误")
		return
	}

	// 查询定时任务
	info, err := service.NewUser(global.DB).GetById(idCondition)
	if err != nil {
		control.FailByMsgf("查询用户信息失败 %s", err.Error())
		return
	}

	if info.Mark == global.SYS_AUTO_CREATE {
		control.FailByMsgf("用户信息【%s-%d】由系统生成, 不允许删除", info.Name, info.ID)
		return
	}

	err = service.NewUser(global.DB).Delete(idCondition)
	if err != nil {
		control.FailByMsgf("删除用户信息失败 %s", err.Error())
		return
	}

	control.Success()
}
