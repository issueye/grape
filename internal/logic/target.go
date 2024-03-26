package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Target struct{}

func (Target) Get(req *repository.QueryTarget) ([]*model.TargetInfo, error) {
	return service.NewTarget().Query(req)
}

func (Target) GetById(id string) (*model.TargetInfo, error) {
	return service.NewTarget().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Target) Modify(id string, req *repository.ModifyTarget) error {
	TargetService := service.NewTarget()
	return TargetService.Modify(id, req)
}

// Modify
// 修改信息 不包含状态
func (Target) ModifyState(id string) error {

	info, err := Target{}.GetById(id)
	if err != nil {
		return err
	}

	state := uint(0)
	if info.State == 0 {
		state = 1
	}

	TargetService := service.NewTarget()
	return TargetService.ModifyState(id, state)
}

// Create
// 创建数据
func (Target) Create(req *repository.CreateTarget) error {
	// 判断端口号在当前系统是否已经被使用
	TargetService := service.NewTarget()
	info, err := TargetService.FindByName(req.Name)
	if err != nil {
		return fmt.Errorf("检查目标地址失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("目标地址[%s]已经添加，请勿重复添加", info.Name)
	}

	// 创建数据
	err = TargetService.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Target) Del(id string) error {
	TargetService := service.NewTarget()

	// 检查使用状态，如果是正在使用则不允许删除
	_, err := TargetService.FindById(id)
	if err != nil {
		return err
	}

	return TargetService.Del(id)
}
