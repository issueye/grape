package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Port struct{}

func (Port) Get(req *repository.QueryPort) ([]*model.PortInfo, error) {
	return service.NewPort().Query(req)
}

func (Port) GetById(id string) (*model.PortInfo, error) {
	return service.NewPort().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Port) Modify(req *repository.ModifyPort) error {
	portServie := service.NewPort()
	return portServie.Modify(req)
}

// ModifyState
// 修改使用状态 返回修改之后的状态
func (Port) ModifyState(id string) (bool, error) {
	portServie := service.NewPort()

	info, err := portServie.FindById(id)
	if err != nil {
		return false, err
	}

	err = portServie.ModifyState(id, !info.State)
	if err != nil {
		return false, err
	}

	return !info.State, nil
}

// Create
// 创建数据
func (Port) Create(req *repository.CreatePort) error {
	// 判断端口号在当前系统是否已经被使用
	portServie := service.NewPort()
	info, err := portServie.FindByPort(req.Port)
	if err != nil {
		return fmt.Errorf("检查端口失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("端口[%d]信息已经创建", req.Port)
	}

	// 创建数据
	err = portServie.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Port) Del(id string) error {
	portServie := service.NewPort()

	// 检查使用状态，如果是正在使用则不允许删除
	pi, err := portServie.FindById(id)
	if err != nil {
		return err
	}

	if pi.State {
		return fmt.Errorf("[%d]端口号正在被使用，不能删除", pi.Port)
	}

	return portServie.Del(id)
}