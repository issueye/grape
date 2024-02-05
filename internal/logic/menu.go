package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Menu struct{}

func (Menu) Delete(id string) error {
	// 查询用户信息
	info, err := service.NewMenu().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户信息失败 %s", err.Error())
	}

	err = service.NewMenu().Delete(info.ID)
	if err != nil {
		return fmt.Errorf("删除用户信息失败 %s", err.Error())
	}

	return nil
}

func (Menu) ModifyState(id string) error {
	// 获取当前定时任务的状态
	info, err := service.NewMenu().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户信息失败 %s", err.Error())
	}

	state := uint(0)
	if info.State == 0 {
		state = 1
	}

	err = service.NewMenu().Status(&repository.StatusMenu{
		ID:    id,
		State: state,
	})
	if err != nil {
		return fmt.Errorf("修改用户信息失败 %s", err.Error())
	}

	return nil
}

func (Menu) Modify(id string, data *repository.ModifyMenu) error {
	return service.NewMenu().Modify(id, data)
}
