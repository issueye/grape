package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Resource struct {
	*service.BaseService
}

func NewResource() *Resource {
	return &Resource{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Resource) Create(data *repository.CreateResource) error {
	info := model.ResourceInfo{}.New()
	info.Name = data.Name
	info.Mark = data.Mark

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Resource) Query(req *repository.QueryResource) ([]*model.ResourceInfo, error) {
	list := make([]*model.ResourceInfo, 0)

	err := s.DataFilter(model.ResourceInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
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
func (s *Resource) Modify(data *repository.ModifyResource) error {
	updateData := make(map[string]any)
	updateData["name"] = data.Name
	updateData["title"] = data.Title
	updateData["folder"] = data.Folder
	updateData["file_name"] = data.FileName
	updateData["mark"] = data.Mark
	return s.Db.Model(&model.ResourceInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Modify
// 修改信息
func (s *Resource) ModifyByMap(id string, datas map[string]any) error {
	return s.Db.Model(&model.ResourceInfo{}).Where("id = ?", id).Updates(datas).Error
}

// Del
// 删除
func (s *Resource) Del(id string) error {
	return s.Db.Model(&model.ResourceInfo{}).Delete("id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Resource) FindById(id string) (*model.ResourceInfo, error) {
	info := new(model.ResourceInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Resource) FindByName(name string) (*model.ResourceInfo, error) {
	info := new(model.ResourceInfo)
	err := s.Db.Model(info).Where("name = ?", name).Find(info).Error
	return info, err
}
