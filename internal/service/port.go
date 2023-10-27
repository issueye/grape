package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Port struct {
	*service.BaseService
}

func NewPort() *Port {
	return &Port{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Port) Create(data *repository.CreatePort) error {
	info := model.PortInfo{}.New()
	info.Port = data.Port
	info.State = false
	info.IsTLS = data.IsTLS
	info.CerId = data.CertId
	info.Mark = data.Mark

	return s.Db.Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Port) Query(req *repository.QueryPort) ([]*model.PortInfo, error) {
	list := make([]*model.PortInfo, 0)

	err := s.DataFilter(model.PortInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Conditon != "" {
			q = q.Where("convert(varchar, port) like ?", fmt.Sprintf("%%%s%%", req.Conditon))
			q = q.Where("cert_code like ?", fmt.Sprintf("%%%s%%", req.Conditon))
			q = q.Where("mark like ?", fmt.Sprintf("%%%s%%", req.Conditon))
		}

		return q, nil
	})

	return list, err
}

// Modify
// 修改信息
func (s *Port) Modify(data *repository.ModifyPort) error {
	updateData := make(map[string]any)
	updateData["is_tls"] = data.IsTLS
	updateData["cert_id"] = data.CertId
	updateData["mark"] = data.Mark
	return s.Db.Model(&model.PortInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Modify
// 修改信息
func (s *Port) ModifyState(id string, state bool) error {
	return s.Db.Model(&model.PortInfo{}).Where("id = ?", id).Update("state", state).Error
}

// Del
// 删除
func (s *Port) Del(id string) error {
	return s.Db.Model(&model.PortInfo{}).Delete("id = ?", id).Error
}

// FindByPort
// 通过端口号查找信息
func (s *Port) FindByPort(port int) (*model.PortInfo, error) {
	info := new(model.PortInfo)
	err := s.Db.Model(info).Where("port = ?", port).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Port) FindById(id string) (*model.PortInfo, error) {
	info := new(model.PortInfo)
	err := s.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}
