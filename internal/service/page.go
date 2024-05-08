package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Page struct {
	*service.BaseService
}

func NewPage() *Page {
	return &Page{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Page) Create(data *repository.CreatePage) error {
	info := model.PageInfo{}.New()
	info.Name = data.Name
	info.PortId = data.PortId
	info.Mark = data.Mark

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Page) Query(req *repository.QueryPage) ([]*model.PageInfo, error) {
	list := make([]*model.PageInfo, 0)

	err := s.DataFilter(model.PageInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Condition != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Condition)).
				Or("mark like ?", fmt.Sprintf("%%%s%%", req.Condition))
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
func (s *Page) Modify(data *repository.ModifyPage) error {
	updateData := make(map[string]any)
	updateData["name"] = data.Name
	updateData["page_path"] = data.PagePath
	updateData["port_id"] = data.PortId
	updateData["mark"] = data.Mark
	return s.Db.Model(&model.PageInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Modify
// 修改信息
func (s *Page) ModifyByMap(id string, datas map[string]any) error {
	return s.Db.Model(&model.PageInfo{}).Where("id = ?", id).Updates(datas).Error
}

// Del
// 删除
func (s *Page) Del(id string) error {
	return s.Db.Model(&model.PageInfo{}).Delete("id = ?", id).Error
}

// Del
// 删除
func (s *Page) DelByPortId(id string) error {
	return s.Db.Model(&model.PageInfo{}).Delete("port_id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Page) FindById(id string) (*model.PageInfo, error) {
	info := new(model.PageInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Page) FindByName(name string, portId string) (*model.PageInfo, error) {
	info := new(model.PageInfo)
	err := s.Db.Model(info).Where("name = ?", name).Where("port_id = ?", portId).Find(info).Error
	return info, err
}
