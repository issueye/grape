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

func (Route) PortCount(id string) (int64, error) {
	return service.NewRule().PortCount(id)
}

// Modify
// 修改信息 不包含状态
func (Route) Modify(req *repository.ModifyRule) error {
	RouteService := service.NewRule()
	return RouteService.Modify(req)
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
	RouteService := service.NewRule()
	info, err := RouteService.FindByName(req.Name, req.PortId)
	if err != nil {
		return fmt.Errorf("检查端口失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("端口[%d]下的匹配规则[%s]已经存在，请勿重复添加", port.Port, info.Name)
	}

	// 创建数据
	err = RouteService.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	err = portService.StepCount(req.PortId, service.ST_RULE, service.STT_PLUS)
	if err != nil {
		return fmt.Errorf("更新页面统计失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Route) Del(id string) error {
	RouteService := service.NewRule()

	// 检查使用状态，如果是正在使用则不允许删除
	info, err := RouteService.FindById(id)
	if err != nil {
		return err
	}

	err = RouteService.Del(id)
	if err != nil {
		return err
	}

	portService := service.NewPort()
	err = portService.StepCount(info.PortId, service.ST_RULE, service.STT_REDUCE)
	return err
}
