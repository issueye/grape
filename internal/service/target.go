package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Target struct {
	*service.BaseService
}

func NewTarget() *Target {
	return &Target{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Target) Create(data *repository.CreateTarget) error {
	info := model.TargetInfo{}.New()
	info.Name = data.Name
	info.Mark = data.Mark

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Target) Query(req *repository.QueryTarget) ([]*model.TargetInfo, error) {
	list := make([]*model.TargetInfo, 0)

	err := s.DataFilter(model.TargetInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Condition != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Condition)).
				Or("mark like ?", fmt.Sprintf("%%%s%%", req.Condition))
		}

		return q, nil
	})

	return list, err
}

// Modify
// 修改信息
func (s *Target) Modify(data *repository.ModifyTarget) error {
	updateData := make(map[string]any)
	updateData["name"] = data.Name
	updateData["mark"] = data.Mark
	return s.Db.Model(&model.TargetInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Del
// 删除
func (s *Target) Del(id string) error {
	return s.Db.Model(&model.TargetInfo{}).Delete("id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Target) FindById(id string) (*model.TargetInfo, error) {
	info := new(model.TargetInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Target) FindByName(name string) (*model.TargetInfo, error) {
	info := new(model.TargetInfo)
	err := s.Db.Model(info).Where("name = ?", name).Find(info).Error
	return info, err
}
