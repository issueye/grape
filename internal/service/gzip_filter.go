package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type GzipFilter struct {
	service.BaseService
}

func (owner *GzipFilter) Self() *GzipFilter {
	return owner
}

func (owner *GzipFilter) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewGzipFilter(args ...service.ServiceContext) *GzipFilter {
	return service.NewServiceSelf(&GzipFilter{}, args...)
}

// Create
// 创建信息
func (s *GzipFilter) Create(data *repository.CreateGzipFilter) error {
	info := model.GzipFilterInfo{}.New()
	info.PortId = data.PortId
	info.MatchContent = data.MatchContent
	info.MatchType = data.MatchType
	info.Mark = data.Mark

	return s.GetDB().Model(info).Create(info).Error
}

// Query
// 查询数据
func (s *GzipFilter) Query(req *repository.QueryGzipFilter) ([]*model.GzipFilterInfo, error) {
	list := make([]*model.GzipFilterInfo, 0)

	err := s.DataFilter(model.GzipFilterInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		qry := db.Order("created_at")
		qry = qry.Where("port_id = ?", req.PortId)

		if req.Condition != "" {
			qry = qry.Where("convert(varchar, port) like ?", fmt.Sprintf("%%%s%%", req.Condition)).
				Or("match_content like ?", fmt.Sprintf("%%%s%%", req.Condition))
		}

		return qry, nil
	})

	return list, err
}

// Modify
// 修改信息
func (s *GzipFilter) Modify(id string, data *repository.ModifyGzipFilter) error {
	updateData := make(map[string]any)
	updateData["match_content"] = data.MatchContent
	updateData["match_type"] = data.MatchType
	updateData["mark"] = data.Mark

	return s.GetDB().Model(&model.GzipFilterInfo{}).Where("id = ?", id).Updates(updateData).Error
}

// Del
// 删除
func (s *GzipFilter) Del(id string) error {
	return s.GetDB().Model(&model.GzipFilterInfo{}).Delete("id = ?", id).Error
}

// FindByPort
// 通过端口号查找信息
func (s *GzipFilter) FindByPort(port int) (*model.GzipFilterInfo, error) {
	info := new(model.GzipFilterInfo)
	err := s.GetDB().Model(info).Where("port = ?", port).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *GzipFilter) FindById(id string) (*model.GzipFilterInfo, error) {
	info := new(model.GzipFilterInfo)
	err := s.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *GzipFilter) FindByContent(portId string, content string) (*model.GzipFilterInfo, error) {
	info := new(model.GzipFilterInfo)
	err := s.GetDB().Model(info).Where("port_id = ?", portId).Where("match_content = ?", content).Find(info).Error
	return info, err
}
