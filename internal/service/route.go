package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Rule struct {
	service.BaseService
}

func (owner *Rule) Self() *Rule {
	return owner
}

func (owner *Rule) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewRule(args ...service.ServiceContext) *Rule {
	return service.NewServiceSelf(&Rule{}, args...)
}

// Create
// 创建信息
func (s *Rule) Create(data *repository.CreateRule) error {
	info := model.RuleInfo{}.New()
	info.Name = data.Name
	info.MatchType = data.MatchType
	info.Method = data.Method
	info.Mark = data.Mark
	info.PortId = data.PortId
	info.TargetId = data.TargetId
	info.TargetRoute = data.TargetRoute

	return s.GetDB().Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Rule) Query(req *repository.QueryRule) ([]*repository.QueryRuleRes, error) {
	list := make([]*repository.QueryRuleRes, 0)

	sqlStr := `SELECT * FROM (
		SELECT a.*, t.name target FROM (
			SELECT r.*, p.port FROM rule_info r LEFT JOIN port_info p ON r.port_id = p.id
			) a LEFT JOIN target_info t ON a.target_id = t.id
		) tb`

	err := s.DataFilter(fmt.Sprintf("(%s)tb", sqlStr), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Condition != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Condition))
			q = q.Where("mark like ?", fmt.Sprintf("%%%s%%", req.Condition))
		}

		if req.PortId != "" {
			q = q.Where("port_id = ?", req.PortId)
		}

		if req.MatchType > 0 {
			q = q.Where("match_type = ?", req.MatchType)
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
	updateData["method"] = data.Method
	updateData["target_id"] = data.TargetId
	updateData["target_route"] = data.TargetRoute
	updateData["port_id"] = data.PortId
	updateData["mark"] = data.Mark
	return s.GetDB().Model(&model.RuleInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Del
// 删除
func (s *Rule) Del(id string) error {
	return s.GetDB().Model(&model.RuleInfo{}).Delete("id = ?", id).Error
}

// Del
// 删除
func (s *Rule) DelByPortId(id string) error {
	return s.GetDB().Model(&model.RuleInfo{}).Delete("port_id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Rule) FindById(id string) (*model.RuleInfo, error) {
	info := new(model.RuleInfo)
	err := s.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Rule) FindLikeName(name string) (bool, error) {
	list := make([]*model.RuleInfo, 0)
	err := s.GetDB().Model(&model.RuleInfo{}).Where("match_type = ?", 1).Where("name like ?", fmt.Sprintf("/%s%%", name)).Find(&list).Error

	if len(list) > 0 {
		return true, err
	} else {
		return false, err
	}
}

// FindById
// 通过ID查找信息
func (s *Rule) FindByName(name string, portId string) (*model.RuleInfo, error) {
	info := new(model.RuleInfo)
	err := s.GetDB().Model(info).Where("name = ?", name).Where("port_id = ?", portId).Find(info).Error
	return info, err
}
