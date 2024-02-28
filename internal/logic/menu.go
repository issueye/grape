package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
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

// 如果菜单信息不存在则创建
func (Menu) CreateMenuNonExistent() {
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "首页", Name: "dashboard", Route: "/home/dashboard", Icon: "", Auth: 0, Level: 0, ParentId: "0", State: 1, Order: 1})

	// 服务地址管理
	serverId := service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "服务管理", Name: "server_manage", Route: "/home/server", Icon: "", Auth: 0, Level: 0, ParentId: "0", State: 1, Order: 1})
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "服务地址管理", Name: "target_manage", Route: "/home/target_manage", Icon: "", Auth: 0, Level: 1, ParentId: serverId, State: 1, Order: 1})
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "页面管理", Name: "page_manage", Route: "/home/page_manage", Icon: "", Auth: 0, Level: 1, ParentId: serverId, State: 1, Order: 2})

	// 系统管理
	systemId := service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "系统管理", Name: "system_manage", Route: "/home/system", Icon: "", Auth: 0, Level: 0, ParentId: "0", State: 1, Order: 2})
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "用户管理", Name: "user_manage", Route: "/home/user_manage", Icon: "", Auth: 0, Level: 1, ParentId: systemId, State: 1, Order: 1})
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "用户组管理", Name: "user_group_manage", Route: "/home/user_group_manage", Icon: "", Auth: 0, Level: 1, ParentId: systemId, State: 1, Order: 2})
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "菜单管理", Name: "menu_manage", Route: "/home/menu_manage", Icon: "", Auth: 0, Level: 1, ParentId: systemId, State: 1, Order: 3})
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "权限组菜单管理", Name: "group_menu_manage", Route: "/home/group_menu_manage", Icon: "", Auth: 0, Level: 1, ParentId: systemId, State: 1, Order: 4})
	service.NewMenu().CreateNoExistent(&model.MenuBase{Title: "参数管理", Name: "param_manage", Route: "/home/param_manage", Icon: "", Auth: 0, Level: 1, ParentId: systemId, State: 1, Order: 5})
}
