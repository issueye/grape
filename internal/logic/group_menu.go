package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type GroupMenu struct{}

func (GroupMenu) Delete(id string) error {
	// 查询用户信息
	info, err := service.NewGroupMenu().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户信息失败 %s", err.Error())
	}

	err = service.NewGroupMenu().Delete(info.ID)
	if err != nil {
		return fmt.Errorf("删除用户信息失败 %s", err.Error())
	}

	return nil
}

func (GroupMenu) ModifyState(id string) error {
	// 获取当前定时任务的状态
	info, err := service.NewGroupMenu().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户信息失败 %s", err.Error())
	}

	state := uint(0)
	if info.State == 0 {
		state = 1
	}

	err = service.NewGroupMenu().Status(&repository.StatusGroupMenu{
		ID:    id,
		State: state,
	})
	if err != nil {
		return fmt.Errorf("修改用户信息失败 %s", err.Error())
	}

	return nil
}

func (GroupMenu) Modify(id string, data *repository.ModifyGroupMenu) error {
	return service.NewGroupMenu().Modify(id, data)
}
