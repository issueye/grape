package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Node struct {
	*service.BaseService
}

func NewNode() *Node {
	return &Node{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Node) Create(data *repository.CreateNode) error {
	info := model.NodeInfo{}.New()
	info.Name = data.Name
	info.PortId = data.PortId
	info.Target = data.Target
	info.NodeType = data.NodeType
	info.PagePath = data.PagePath
	info.Mark = data.Mark

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Node) Query(req *repository.QueryNode) ([]*model.NodeInfo, error) {
	list := make([]*model.NodeInfo, 0)

	err := s.DataFilter(model.NodeInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Conditon != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Conditon))
			q = q.Where("page_path like ?", fmt.Sprintf("%%%s%%", req.Conditon))
			q = q.Where("target like ?", fmt.Sprintf("%%%s%%", req.Conditon))
			q = q.Where("mark like ?", fmt.Sprintf("%%%s%%", req.Conditon))
		}

		if req.PortId != "" {
			q = q.Where("port_id = ?", req.PortId)
		}

		return q, nil
	})

	return list, err
}

// Modify
// 修改信息
func (s *Node) Modify(data *repository.ModifyNode) error {
	updateData := make(map[string]any)
	updateData["name"] = data.Name
	updateData["node_type"] = data.NodeType
	updateData["page_path"] = data.PagePath
	updateData["port_id"] = data.PortId
	return s.Db.Model(&model.NodeInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Del
// 删除
func (s *Node) Del(id string) error {
	return s.Db.Model(&model.NodeInfo{}).Delete("id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Node) FindById(id string) (*model.NodeInfo, error) {
	info := new(model.NodeInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Node) FindByName(name string) (*model.NodeInfo, error) {
	info := new(model.NodeInfo)
	err := s.Db.Model(info).Where("name = ?", name).Find(info).Error
	return info, err
}
