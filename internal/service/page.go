package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Page struct {
	service.BaseService
}

func (owner *Page) Self() *Page {
	return owner
}

func (owner *Page) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewPage(args ...service.ServiceContext) *Page {
	return service.NewServiceSelf(&Page{}, args...)
}

// Create
// 创建信息
func (s *Page) Create(data *repository.CreatePage) error {
	info := model.PageInfo{}.New()
	info.Name = data.Name
	info.Title = data.Title
	info.PortId = data.PortId
	info.Version = data.Version
	info.ProductCode = data.ProductCode
	info.Mark = data.Mark

	err := s.GetDB().Model(info).Create(info).Error
	if err != nil {
		return err
	}

	return nil
}

// Create
// 创建信息
func (s *Page) CreatePageVersion(data *model.PageVersionBase) error {
	version := model.PageVersionInfo{}.New()
	version.Copy(data)

	return s.GetDB().Create(version).Error
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

// Query
// 查询数据
func (s *Page) QueryVersion(req *repository.QueryPageVersion) ([]*model.PageVersionInfo, error) {
	list := make([]*model.PageVersionInfo, 0)

	err := s.DataFilter(model.PageVersionInfo{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		q := db.Order("created_at")

		if req.Condition != "" {
			q = q.Where("name like ?", fmt.Sprintf("%%%s%%", req.Condition)).
				Or("mark like ?", fmt.Sprintf("%%%s%%", req.Condition))
		}

		if req.PortId != "" {
			q = q.Where("port_id = ?", req.PortId)
		}

		if req.ProductCode != "" {
			q = q.Where("product_code = ?", req.ProductCode)
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
	updateData["thumbnail"] = data.Thumbnail
	updateData["version"] = data.Version
	updateData["port_id"] = data.PortId
	updateData["mark"] = data.Mark
	return s.GetDB().Model(&model.PageInfo{}).Where("id = ?", data.ID).Updates(updateData).Error
}

// Modify
// 修改信息
func (s *Page) ModifyByMap(id string, datas map[string]any) error {
	return s.GetDB().Model(&model.PageInfo{}).Where("id = ?", id).Updates(datas).Error
}

// Del
// 删除
func (s *Page) Del(id string) error {
	return s.GetDB().Where("id = ?", id).Delete(&model.PageInfo{}).Error
}

// Del
// 删除
func (s *Page) DelAllVersion(productCode string) error {
	return s.GetDB().Where("product_code = ?", productCode).Delete(&model.PageVersionInfo{}).Error
}

// Del
// 删除
func (s *Page) DelByPortId(id string) error {
	return s.GetDB().Model(&model.PageInfo{}).Delete("port_id = ?", id).Error
}

// FindById
// 通过ID查找信息
func (s *Page) FindById(id string) (*model.PageInfo, error) {
	info := new(model.PageInfo)
	err := s.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Page) FindByName(name string, portId string) (*model.PageInfo, error) {
	info := new(model.PageInfo)
	err := s.GetDB().Model(info).Where("name = ?", name).Where("port_id = ?", portId).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Page) FindByProductCode(portId, productCode string) (*model.PageInfo, error) {
	info := new(model.PageInfo)
	err := s.GetDB().Model(info).Where("port_id = ?", portId).Where("product_code = ?", productCode).Find(info).Error
	return info, err
}

// FindById
// 通过ID查找信息
func (s *Page) FindByVersion(portId string, productCode string, version string) (*model.PageVersionInfo, error) {
	info := new(model.PageVersionInfo)
	err := s.GetDB().Model(info).Where("port_id = ?", portId).Where("product_code = ?", productCode).Where("version = ?", version).Find(info).Error
	return info, err
}
