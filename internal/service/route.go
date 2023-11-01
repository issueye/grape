package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Rule struct {
	*service.BaseService
}

func NewRule() *Rule {
	return &Rule{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Rule) Create(data *repository.CreateRule) error {
	info := model.RuleInfo{}.New()
	info.Name = data.Name
	info.MatchType = data.MatchType
	info.Mark = data.Mark
	info.NodeId = data.NodeId
	info.PortId = data.PortId

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Rule) Query(req *repository.QueryRule) ([]*repository.QueryRuleRes, error) {
	list := make([]*repository.QueryRuleRes, 0)

	sqlStr := "select rule.*, port from rule_info rule, port_info port where rule.port_id = port.id"

	err := s.DataFilter(fmt.Sprintf("(%s)tb", sqlStr), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Conditon != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Conditon))
			q = q.Where("mark like ?", fmt.Sprintf("%%%s%%", req.Conditon))
		}

		if req.PortId != "" {
			q = q.Where("port_id = ?", req.PortId)
		}

		if req.NodeId != "" {
			q = q.Where("node_id = ?", req.NodeId)
		}

		return q, nil
	})

	return list, err
}

// Modify
// 修改信息
func (s *Rule) Modify(data *repository.ModifyRule) error {
	updateData := make(map[string]any)
	updateData["name"] = data.Name
	updateData["match_type"] = data.MatchType
	updateData["target"] = data.Target
	updateData["node_id"] = data.NodeId
	updateData["port_id"] = data.PortId
	updateData["mark"] = data.Mark
	return s.Db.Model(&model.RuleInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Del
// 删除
func (s *Rule) Del(id string) error {
	return s.Db.Model(&model.RuleInfo{}).Delete("id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Rule) FindById(id string) (*model.RuleInfo, error) {
	info := new(model.RuleInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Rule) FindByName(name, portId string) (*model.RuleInfo, error) {
	info := new(model.RuleInfo)
	err := s.Db.Model(info).Where("name = ?", name).Where("port_id = ?", portId).Find(info).Error
	return info, err
}
