package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Cert struct {
	*service.BaseService
}

func NewCert() *Cert {
	return &Cert{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Cert) Create(data *repository.CreateCert) error {
	info := model.CertInfo{}.New()
	info.Name = data.Name
	info.Public = data.Public
	info.Private = data.Private
	info.Mark = data.Mark

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Cert) Query(req *repository.QueryCert) ([]*model.CertInfo, error) {
	list := make([]*model.CertInfo, 0)

	err := s.DataFilter(model.CertInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Condition != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Condition)).
				Or("mark like ?", fmt.Sprintf("%%%s%%", req.Condition)).
				Or("public like ?", fmt.Sprintf("%%%s%%", req.Condition)).
				Or("private like ?", fmt.Sprintf("%%%s%%", req.Condition))
		}

		return q, nil
	})

	return list, err
}

// Modify
// 修改信息
func (s *Cert) Modify(data *repository.ModifyCert) error {
	updateData := make(map[string]any)
	updateData["name"] = data.Name
	updateData["public"] = data.Public
	updateData["private"] = data.Private
	updateData["mark"] = data.Mark
	return s.Db.Model(&model.CertInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Del
// 删除
func (s *Cert) Del(id string) error {
	return s.Db.Model(&model.CertInfo{}).Delete("id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Cert) FindById(id string) (*model.CertInfo, error) {
	info := new(model.CertInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Cert) FindByName(name string) (*model.CertInfo, error) {
	info := new(model.CertInfo)
	err := s.Db.Model(info).Where("name = ?", name).Find(info).Error
	return info, err
}
