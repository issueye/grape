package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
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
func (Port) Modify(id string, req *repository.ModifyPort) error {
	portService := service.NewPort()
	return portService.Modify(id, req)
}

// Modify
// 修改信息 不包含状态
func (Port) ModifyByMap(id string, data map[string]any) error {
	portService := service.NewPort()
	return portService.ModifyByMap(id, data)
}

// ModifyState
// 修改使用状态 返回修改之后的状态
func (Port) ModifyState(id string) (bool, error) {
	portService := service.NewPort()

	info, err := portService.FindById(id)
	if err != nil {
		return false, err
	}

	err = portService.ModifyState(id, !info.State)
	if err != nil {
		return false, err
	}

	return !info.State, nil
}

// 刷新端口统计
func (Port) RefreshStatistics() error {
	list, err := Port{}.Get(&repository.QueryPort{})
	if err != nil {
		return err
	}

	for _, portInfo := range list {
		// 更新页面
		count, err := Page{}.PortCount(portInfo.ID)
		if err != nil {
			continue
		}
		portInfo.PageCount = int(count)

		count, err = Route{}.PortCount(portInfo.ID)
		if err != nil {
			continue
		}
		portInfo.RuleCount = int(count)

		count, err = GzipFilter{}.PortCount(portInfo.ID)
		if err != nil {
			continue
		}
		portInfo.GzipFilterCount = int(count)

		// 更新数据
		err = Port{}.ModifyByMap(portInfo.ID, map[string]any{
			"page_count":        portInfo.PageCount,
			"rule_count":        portInfo.RuleCount,
			"gzip_filter_count": portInfo.GzipFilterCount,
		})

		if err != nil {
			global.Log.Errorf("端口号[%d]更新统计信息失败 %s", portInfo.Port, err.Error())
			continue
		}
	}

	return nil
}

func (Port) Notice(id string) error {
	info, err := Port{}.GetById(id)
	if err != nil {
		return err
	}

	// 处理状态
	// fmt.Println("当前端口状态", info.State)
	ch := &global.Port{PortInfo: *info}

	if info.State {
		ch.Action = global.AT_START
	} else {
		ch.Action = global.AT_STOP
	}

	global.PortChan <- ch

	return nil
}

// ModifyState
// 修改使用状态 返回修改之后的状态
func (Port) Stop(id string) error {
	return service.NewPort().ModifyState(id, false)
}

// ModifyState
// 修改使用状态 返回修改之后的状态
func (Port) Start(id string) error {
	return service.NewPort().ModifyState(id, true)
}

// ModifyState
// 修改使用状态 返回修改之后的状态
func (Port) ModifyGzip(id string) error {
	info, err := Port{}.GetById(id)
	if err != nil {
		return fmt.Errorf("查询端口信息失败 %s", err.Error())
	}

	use := !info.UseGzip
	return service.NewPort().ModifyGzip(id, use)
}

// Create
// 创建数据
func (Port) Create(req *repository.CreatePort) error {
	// 判断端口号在当前系统是否已经被使用
	portService := service.NewPort()
	info, err := portService.FindByPort(req.Port)
	if err != nil {
		return fmt.Errorf("检查端口失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("端口[%d]信息已经创建", req.Port)
	}

	// 创建数据
	err = portService.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Port) Del(id string) error {
	portService := service.NewPort()

	// 检查使用状态，如果是正在使用则不允许删除
	pi, err := portService.FindById(id)
	if err != nil {
		return err
	}

	if pi.State {
		return fmt.Errorf("[%d]端口号正在被使用，不能删除", pi.Port)
	}

	err = portService.Del(id)
	if err != nil {
		return fmt.Errorf("删除端口号[%d]失败 %s", pi.Port, err.Error())
	}

	// 删除匹配规则
	err = service.NewRule().DelByPortId(id)
	if err != nil {
		return fmt.Errorf("删除端口号[%d]下的匹配规则失败 %s", pi.Port, err.Error())
	}

	// 删除节点
	service.NewPage().DelByPortId(id)
	if err != nil {
		return fmt.Errorf("删除端口号[%d]下的节点失败 %s", pi.Port, err.Error())
	}

	return nil
}
