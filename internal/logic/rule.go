package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Route struct{}

func (Route) Get(req *repository.QueryRule) ([]*repository.QueryRuleRes, error) {
	return service.NewRule().Query(req)
}

func (Route) GetById(id string) (*model.RuleInfo, error) {
	return service.NewRule().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Route) Modify(req *repository.ModifyRule) error {
	RouteServie := service.NewRule()
	return RouteServie.Modify(req)
}

// Create
// 创建数据
func (Route) Create(req *repository.CreateRule) error {
	portService := service.NewPort()
	port, err := portService.FindById(req.PortId)
	if err != nil {
		return fmt.Errorf("获取端口信息失败 %s", err.Error())
	}

	// 判断端口号在当前系统是否已经被使用
	RouteServie := service.NewRule()
	info, err := RouteServie.FindByName(req.Name, req.PortId)
	if err != nil {
		return fmt.Errorf("检查端口失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("端口[%d]下的匹配规则[%s]已经存在，请勿重复添加", port.Port, info.Name)
	}

	// 创建数据
	err = RouteServie.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Route) Del(id string) error {
	RouteServie := service.NewRule()

	// 检查使用状态，如果是正在使用则不允许删除
	_, err := RouteServie.FindById(id)
	if err != nil {
		return err
	}

	return RouteServie.Del(id)
}
