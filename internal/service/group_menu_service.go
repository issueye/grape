package service

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"gorm.io/gorm"
)

type GroupMenu struct {
	*service.BaseService
}

func NewGroupMenu() *GroupMenu {
	return &GroupMenu{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建用户组菜单信息
func (GroupMenu *GroupMenu) Create(data *repository.CreateGroupMenu) error {
	info := model.GroupMenu{}.New()
	info.Copy(&data.GroupMenuBase)
	info.State = 1
	return GroupMenu.Db.Create(info).Error
}

// Create
// 创建用户组菜单信息
func (GroupMenu *GroupMenu) BatchCreate(datas *[]*model.GroupMenu) error {
	return GroupMenu.Db.Create(datas).Error
}

// GetById
// 根据用户组菜单ID查找用户组菜单信息
func (GroupMenu *GroupMenu) GetById(id string) (*model.GroupMenu, error) {
	info := new(model.GroupMenu)
	err := GroupMenu.Db.Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// Modify
// 修改用户组菜单信息
func (GroupMenu *GroupMenu) Modify(id string, info *repository.ModifyGroupMenu) error {
	m := make(map[string]any)
	m["name"] = info.Name
	m["route"] = info.Route
	m["title"] = info.Title
	m["icon"] = info.Icon
	m["auth"] = info.Auth
	m["level"] = info.Level
	m["parent_id"] = info.ParentId

	return GroupMenu.Db.Model(&model.GroupMenu{}).Where("id = ?", id).Updates(m).Error
}

// Status
// 修改用户组菜单信息
func (GroupMenu *GroupMenu) Status(info *repository.StatusGroupMenu) error {
	return GroupMenu.Db.
		Model(&model.GroupMenu{}).
		Where("id = ?", info.ID).
		Update("state", info.State).
		Error
}

// Delete
// 删除用户组菜单信息
func (GroupMenu *GroupMenu) Delete(id string) error {
	return GroupMenu.Db.Where("id = ?", id).Delete(&model.GroupMenu{}).Error
}

// List
// 获取用户组菜单列表
func (GroupMenu *GroupMenu) List(info *repository.QueryGroupMenu) ([]*model.GroupMenu, error) {
	GroupMenuInfo := new(model.GroupMenu)
	list := make([]*model.GroupMenu, 0)
	err := GroupMenu.DataFilter(GroupMenuInfo.TableName(), info, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("level").Order("[order]")

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
