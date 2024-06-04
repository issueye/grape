package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Port struct {
	service.BaseService
}

func (owner *Port) Self() *Port {
	return owner
}

func (owner *Port) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewPort(args ...service.ServiceContext) *Port {
	return service.NewServiceSelf(&Port{}, args...)
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

	return s.GetDB().Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *Port) Query(req *repository.QueryPort) ([]*model.PortInfo, error) {
	list := make([]*model.PortInfo, 0)

	err := s.DataFilter(model.PortInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Condition != "" {
			q = q.Where("CAST(port AS TEXT) like ?", fmt.Sprintf("%%%s%%", req.Condition))
			q = q.Or("mark like ?", fmt.Sprintf("%%%s%%", req.Condition))
		}

		return q, nil
	})

	return list, err
}

// Modify
// 修改信息
func (s *Port) Modify(id string, data *repository.ModifyPort) error {
	updateData := make(map[string]any)
	updateData["is_tls"] = data.IsTLS
	updateData["cert_id"] = data.CertId
	updateData["mark"] = data.Mark
	updateData["use_gzip"] = data.UseGzip
	return s.GetDB().Model(&model.PortInfo{}).Where("id = ?", id).Updates(updateData).Error
}

type StatisticsType int

const (
	ST_PAGE StatisticsType = iota
	ST_RULE
	ST_GZIP_FILTER
)

type StepType int

const (
	STT_PLUS StepType = iota
	STT_REDUCE
)

// Modify
// 修改信息
func (s *Port) StepCount(id string, t StatisticsType, st StepType) error {

	stepStr := " + 1"
	if st == STT_REDUCE {
		stepStr = " - 1"
	}

	switch t {
	case ST_PAGE:
		return s.GetDB().Raw(fmt.Sprintf("update ? set page_count = page_count %s where id = ?", stepStr), model.PageInfo{}.TableName(), id).Error
	case ST_RULE:
		return s.GetDB().Raw(fmt.Sprintf("update ? set rule_count = rule_count %s where id = ?", stepStr), model.PageInfo{}.TableName(), id).Error
	case ST_GZIP_FILTER:
		return s.GetDB().Raw(fmt.Sprintf("update ? set gzip_filter_count = gzip_filter_count %s where id = ?", stepStr), model.PageInfo{}.TableName(), id).Error
	}

	return nil
}

// Modify
// 修改信息
func (s *Port) ModifyByMap(id string, data map[string]any) error {
	return s.GetDB().Model(&model.PortInfo{}).Where("id = ?", id).Updates(data).Error
}

// Modify
// 修改信息
func (s *Port) ModifyState(id string, state bool) error {
	return s.GetDB().Model(&model.PortInfo{}).Where("id = ?", id).Update("state", state).Error
}

// Modify
// 修改信息
func (s *Port) ModifyGzip(id string, use bool) error {
	return s.GetDB().Model(&model.PortInfo{}).Where("id = ?", id).Update("use_gzip", use).Error
}

// Del
// 删除
func (s *Port) Del(id string) error {
	return s.GetDB().Model(&model.PortInfo{}).Delete("id = ?", id).Error
}

// FindByPort
// 通过端口号查找信息
func (s *Port) FindByPort(port int) (*model.PortInfo, error) {
	info := new(model.PortInfo)
	err := s.GetDB().Model(info).Where("port = ?", port).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Port) FindById(id string) (*model.PortInfo, error) {
	info := new(model.PortInfo)
	err := s.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}
