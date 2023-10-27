package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Route struct{}

func (Route) Get(req *repository.QueryRoute) ([]*model.RouteInfo, error) {
	return service.NewRoute().Query(req)
}

func (Route) GetById(id string) (*model.RouteInfo, error) {
	return service.NewRoute().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Route) Modify(req *repository.ModifyRoute) error {
	RouteServie := service.NewRoute()
	return RouteServie.Modify(req)
}

// Create
// 创建数据
func (Route) Create(req *repository.CreateRoute) error {
	// 判断端口号在当前系统是否已经被使用
	RouteServie := service.NewRoute()
	_, err := RouteServie.FindByName(req.Name)
	if err != nil {
		return fmt.Errorf("检查端口失败 %s", err.Error())
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
	RouteServie := service.NewRoute()

	// 检查使用状态，如果是正在使用则不允许删除
	_, err := RouteServie.FindById(id)
	if err != nil {
		return err
	}

	return RouteServie.Del(id)
}
