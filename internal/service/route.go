package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Route struct {
	*service.BaseService
}

func NewRoute() *Route {
	return &Route{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Route) Create(data *repository.CreateRoute) error {
	info := model.RouteInfo{}.New()
	info.Name = data.Name
	info.MatchType = data.MatchType
	info.Mark = data.Mark

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Route) Query(req *repository.QueryRoute) ([]*model.RouteInfo, error) {
	list := make([]*model.RouteInfo, 0)

	err := s.DataFilter(model.RouteInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Conditon != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Conditon))
			q = q.Where("mark like ?", fmt.Sprintf("%%%s%%", req.Conditon))
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
func (s *Route) Modify(data *repository.ModifyRoute) error {
	updateData := make(map[string]any)
	updateData["name"] = data.Name
	updateData["match_type"] = data.MatchType
	updateData["target"] = data.Target
	updateData["node_id"] = data.NodeId
	updateData["mark"] = data.Mark
	return s.Db.Model(&model.RouteInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Del
// 删除
func (s *Route) Del(id string) error {
	return s.Db.Model(&model.RouteInfo{}).Delete("id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Route) FindById(id string) (*model.RouteInfo, error) {
	info := new(model.RouteInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Route) FindByName(name string) (*model.RouteInfo, error) {
	info := new(model.RouteInfo)
	err := s.Db.Model(info).Where("name = ?", name).Find(info).Error
	return info, err
}
