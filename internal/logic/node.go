package logic

import (
	"fmt"
	"strings"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Page struct{}

func (Page) Get(req *repository.QueryPage) ([]*model.PageInfo, error) {
	return service.NewPage().Query(req)
}

func (Page) GetById(id string) (*model.PageInfo, error) {
	return service.NewPage().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Page) Modify(req *repository.ModifyPage) error {
	PageService := service.NewPage()
	return PageService.Modify(req)
}

// Modify
// 修改信息 不包含状态
func (Page) CheckData(portId string, args ...any) error {
	PageService := service.NewPage()
	list, err := PageService.Query(&repository.QueryPage{
		PortId: portId,
	})

	if err != nil {
		return err
	}

	// 对比名字
	if len(args) > 0 {
		name := args[0].(string)
		for _, Page := range list {
			if strings.HasPrefix(name, fmt.Sprintf("/%s", Page.Name)) {
				return fmt.Errorf("[GIN匹配]类型的匹配规则与页面[%s]冲突，请修改为[MUX匹配]", Page.Name)
			}
		}
	}

	for _, Page := range list {
		// 节点里面去查询所有的节点
		ok, err := service.NewRule().FindLikeName(Page.Name)
		if err != nil {
			return err
		}

		// 查询到有相似的直接返回错误
		if ok {
			return fmt.Errorf("[GIN匹配]类型的匹配规则中发现了与页面[%s]冲突的路由，请修改为[MUX匹配]", Page.Name)
		}
	}

	return nil
}

// Modify
// 修改信息 不包含状态
func (Page) ModifyByMap(id string, datas map[string]any) error {
	PageService := service.NewPage()
	return PageService.ModifyByMap(id, datas)
}

// Create
// 创建数据
func (Page) Create(req *repository.CreatePage) error {
	// 判断端口号在当前系统是否已经被使用
	PageService := service.NewPage()
	info, err := PageService.FindByName(req.Name, req.PortId)
	if err != nil {
		return fmt.Errorf("检查节点失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("节点[%s]已经添加，请勿重复添加", info.Name)
	}

	// 创建数据
	err = PageService.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Page) Del(id string) error {
	PageService := service.NewPage()

	// 检查使用状态，如果是正在使用则不允许删除
	_, err := PageService.FindById(id)
	if err != nil {
		return err
	}

	return PageService.Del(id)
}
