package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type GzipFilter struct{}

func (GzipFilter) Get(req *repository.QueryGzipFilter) ([]*model.GzipFilterInfo, error) {
	return service.NewGzipFilter().Query(req)
}

func (GzipFilter) GetById(id string) (*model.GzipFilterInfo, error) {
	return service.NewGzipFilter().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (GzipFilter) Modify(id string, req *repository.ModifyGzipFilter) error {
	gfService := service.NewGzipFilter()
	return gfService.Modify(id, req)
}

// Create
// 创建数据
func (GzipFilter) Create(req *repository.CreateGzipFilter) error {
	// 判断端口号在当前系统是否已经被使用
	gfService := service.NewGzipFilter()
	info, err := gfService.FindByContent(req.PortId, req.MatchContent)
	if err != nil {
		return fmt.Errorf("检查端口失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("[%s]信息已经创建", req.MatchContent)
	}

	// 创建数据
	err = gfService.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (GzipFilter) Del(id string) error {
	gfService := service.NewGzipFilter()
	portService := service.NewPort()

	// 检查使用状态，如果是正在使用则不允许删除
	pi, err := gfService.FindById(id)
	if err != nil {
		return err
	}

	portInfo, err := portService.FindById(pi.PortId)
	if err != nil {
		return err
	}

	if portInfo.State {
		return fmt.Errorf("[%d]端口号正在被使用，不能删除", portInfo.Port)
	}

	err = gfService.Del(id)
	if err != nil {
		return fmt.Errorf("删除过滤信息[%s]失败 %s", pi.MatchContent, err.Error())
	}

	// 删除匹配规则
	err = gfService.Del(id)
	if err != nil {
		return fmt.Errorf("删除过滤信息[%s]失败 %s", pi.MatchContent, err.Error())
	}

	return nil
}
