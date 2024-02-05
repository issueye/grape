package logic

import (
	"fmt"
	"strings"

	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type User struct{}

func (User) Delete(id string) error {
	// 查询用户信息
	info, err := service.NewUser().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户信息失败 %s", err.Error())
	}

	if info.Mark == global.SYS_AUTO_CREATE {
		return fmt.Errorf("用户信息【%s-%s】由系统生成, 不允许删除", info.Name, info.ID)
	}

	err = service.NewUser().Delete(id)
	if err != nil {
		return fmt.Errorf("删除用户信息失败 %s", err.Error())
	}

	return nil
}

func (User) ModifyState(id string) error {
	// 获取当前定时任务的状态
	info, err := service.NewUser().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户信息失败 %s", err.Error())
	}

	state := uint(0)
	if info.State == 0 {
		state = 1
	}

	err = service.NewUser().Status(&repository.StatusUser{
		ID:    id,
		State: state,
	})
	if err != nil {
		return fmt.Errorf("修改用户信息失败 %s", err.Error())
	}

	return nil
}

func (User) Modify(id string, data *repository.ModifyUser) error {
	// 查询用户信息
	info, err := service.NewUser().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户信息失败 %s", err.Error())
	}

	// 系统任务不允许修改表述信息
	if info.Mark == global.SYS_AUTO_CREATE {
		// 判断描述是否被修改
		if !strings.EqualFold(info.Mark, data.Mark) {
			return fmt.Errorf("【%s-%s】由系统生成, 不允许修改描述信息", info.Name, info.ID)
		}
	}

	err = service.NewUser().Modify(data)
	if err != nil {
		return fmt.Errorf("修改用户信息失败 %s", err.Error())
	}

	return nil
}
