package logic

import (
	"fmt"
	"strings"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Node struct{}

func (Node) Get(req *repository.QueryNode) ([]*model.NodeInfo, error) {
	return service.NewNode().Query(req)
}

func (Node) GetById(id string) (*model.NodeInfo, error) {
	return service.NewNode().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Node) Modify(req *repository.ModifyNode) error {
	NodeService := service.NewNode()
	return NodeService.Modify(req)
}

// Modify
// 修改信息 不包含状态
func (Node) CheckData(portId string, args ...any) error {
	NodeService := service.NewNode()
	list, err := NodeService.Query(&repository.QueryNode{
		PortId: portId,
	})

	if err != nil {
		return err
	}

	// 对比名字
	if len(args) > 0 {
		name := args[0].(string)
		for _, node := range list {
			if strings.HasPrefix(name, fmt.Sprintf("/%s", node.Name)) {
				return fmt.Errorf("[GIN匹配]类型的匹配规则与页面[%s]冲突，请修改为[MUX匹配]", node.Name)
			}
		}
	}

	for _, node := range list {
		// 节点里面去查询所有的节点
		ok, err := service.NewRule().FindLikeName(node.Name)
		if err != nil {
			return err
		}

		// 查询到有相似的直接返回错误
		if ok {
			return fmt.Errorf("[GIN匹配]类型的匹配规则中发现了与页面[%s]冲突的路由，请修改为[MUX匹配]", node.Name)
		}
	}

	return nil
}

// Modify
// 修改信息 不包含状态
func (Node) ModifyByMap(id string, datas map[string]any) error {
	NodeService := service.NewNode()
	return NodeService.ModifyByMap(id, datas)
}

// Create
// 创建数据
func (Node) Create(req *repository.CreateNode) error {
	// 判断端口号在当前系统是否已经被使用
	NodeService := service.NewNode()
	info, err := NodeService.FindByName(req.Name, req.PortId)
	if err != nil {
		return fmt.Errorf("检查节点失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("节点[%s]已经添加，请勿重复添加", info.Name)
	}

	// 创建数据
	err = NodeService.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Node) Del(id string) error {
	NodeService := service.NewNode()

	// 检查使用状态，如果是正在使用则不允许删除
	_, err := NodeService.FindById(id)
	if err != nil {
		return err
	}

	return NodeService.Del(id)
}
