package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type UserGroup struct{}

func (UserGroup) Delete(id string) error {
	// 查询是否存在
	info, err := service.NewUserGroup().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户组信息失败 %s", err.Error())
	}

	if info.Mark == global.SYS_AUTO_CREATE {
		return fmt.Errorf("信息【%s-%s】由系统生成, 不允许删除", info.Name, info.ID)
	}

	err = service.NewUserGroup().Delete(info.ID)
	if err != nil {
		return fmt.Errorf("删除用户组信息失败 %s", err.Error())
	}

	return nil
}

func (UserGroup) ModifyState(id string) error {
	// 获取当前定时任务的状态
	info, err := service.NewUserGroup().GetById(id)
	if err != nil {
		return fmt.Errorf("查询用户组信息失败 %s", err.Error())
	}

	state := uint(0)
	if info.State == 0 {
		state = 1
	}

	err = service.NewUserGroup().Status(&repository.StatusUserGroup{
		ID:    id,
		State: state,
	})

	if err != nil {
		return fmt.Errorf("修改用户组状态失败 %s", err.Error())
	}

	return nil
}

// 检查管理员组中的菜单是否完全，如果不完全则补全
func (UserGroup) CreateAdminNonExistent() error {
	userGroupMenus, err := service.NewGroupMenu().List(&repository.QueryGroupMenu{GroupId: global.AdminGroupId})
	if err != nil {
		return err
	}

	findMenu := func(route string) *model.GroupMenu {
		for _, element := range userGroupMenus {
			if element.Route == route {
				return element
			}
		}

		return nil
	}

	menus, err := service.NewMenu().List(&repository.QueryMenu{Level: -1})
	if err != nil {
		return err
	}

	insertData := make([]*model.GroupMenu, 0)

	for _, element := range menus {
		menu := findMenu(element.Route)
		if menu == nil {
			data := model.GroupMenu{}.New()
			data.Copy2(&element.MenuBase)
			data.MenuId = element.ID
			data.GroupId = global.AdminGroupId
			insertData = append(insertData, data)
		}
	}

	if len(insertData) > 0 {
		return service.NewGroupMenu().BatchCreate(&insertData)
	}

	return nil
}
