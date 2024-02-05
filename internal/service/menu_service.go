package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type Menu struct {
	*service.BaseService
}

func NewMenu() *Menu {
	return &Menu{
		BaseService: service.NewBaseService(global.DB),
	}
}

// CreateAdminNonExistent
// 创建管理员用户，如果不存在
func (Menu *Menu) CreateAdminNonExistent() error {
	// isHave := int64(0)
	// err := Menu.Db.Model(&model.Menu{}).Where("account = ?", global.AdminName).Where("id = ?", global.AdminId).Count(&isHave).Error
	// if err != nil {
	// 	return err
	// }

	// if isHave == 0 {
	// 	info := new(model.Menu)
	// 	info.ID = global.AdminId
	// 	info.Name =

	// 	return Menu.Db.Create(info).Error
	// } else {
	// 	return nil
	// }

	return nil
}

// Create
// 创建用户信息
func (Menu *Menu) Create(data *repository.CreateMenu) error {
	info := model.Menu{}.New()
	info.Copy(&data.MenuBase)

	return Menu.Db.Create(info).Error
}

// GetById
// 根据用户ID查找用户信息
func (Menu *Menu) GetById(id string) (*model.Menu, error) {
	info := new(model.Menu)
	err := Menu.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// Modify
// 修改用户信息
func (Menu *Menu) Modify(id string, info *repository.ModifyMenu) error {
	m := make(map[string]any)
	m["name"] = info.Name
	m["route"] = info.Route
	m["title"] = info.Title
	m["icon"] = info.Icon
	m["auth"] = info.Auth
	m["level"] = info.Level
	m["parent_id"] = info.ParentId

	return Menu.Db.Model(&model.Menu{}).Where("id = ?", id).Updates(m).Error
}

// Status
// 修改用户信息
func (Menu *Menu) Status(info *repository.StatusMenu) error {
	return Menu.Db.
		Model(&model.Menu{}).
		Where("id = ?", info.ID).
		Update("state", info.State).
		Error
}

// Delete
// 删除用户信息
func (Menu *Menu) Delete(id string) error {
	return Menu.Db.Where("id = ?", id).Delete(&model.Menu{}).Error
}

// List
// 获取用户列表
func (Menu *Menu) List(info *repository.QueryMenu) ([]*model.Menu, error) {
	MenuInfo := new(model.Menu)
	list := make([]*model.Menu, 0)
	err := Menu.DataFilter(MenuInfo.TableName(), info, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("id")

		// 通用统一条件
		if info.Condition != "" {
			query = query.
				Where("name like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("title like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("route like ?", fmt.Sprintf("%%%s%%", info.Condition))
		}

		return query, nil
	})
	return list, err
}
