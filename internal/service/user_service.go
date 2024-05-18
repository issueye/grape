package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/config"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	service.BaseService
}

func (owner *User) Self() *User {
	return owner
}

func (owner *User) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewUser(args ...service.ServiceContext) *User {
	return service.NewServiceSelf(&User{}, args...)
}

// FindUser
// 查找用户
func (user *User) FindUser(info *repository.Login) (*model.UserInfo, error) {
	query := user.GetDB().Model(&model.UserInfo{}).Order("id")
	query = query.Where("account = ?", info.Account)

	// 判断是否需要验证密码
	r := config.GetParam("SERVER-MODE", "release")
	fmt.Printf("当前运行模式：%s", r.String())
	if strings.EqualFold(strings.ToLower(r.String()), "release") {
		query = query.Where("password = ?", info.Password)
	}

	data := new(model.UserInfo)
	err := query.Find(data).Error
	return data, err
}

// CreateAdminNonExistent
// 创建管理员用户，如果不存在
func (user *User) CreateAdminNonExistent() error {
	isHave := int64(0)
	err := user.GetDB().Model(&model.UserInfo{}).Where("account = ?", global.AdminName).Where("id = ?", global.AdminId).Count(&isHave).Error
	if err != nil {
		return err
	}

	if isHave == 0 {
		info := new(model.UserInfo)
		info.ID = global.AdminId
		info.GroupId = global.AdminGroupId
		info.Account = global.AdminAccount
		info.Name = global.AdminName
		info.Mark = "系统自动生成的管理员数据"
		info.Password = "123456"
		info.State = 1
		info.Sys = 1

		return user.GetDB().Create(info).Error
	} else {
		return nil
	}
}

// Create
// 创建用户信息
func (user *User) Create(data *repository.CreateUser) error {
	info := new(model.UserInfo)
	info.ID = strconv.FormatInt(utils.GenID(), 10)
	info.Account = data.Account
	info.Name = data.Name
	info.Password = data.Password
	info.Mark = data.Mark
	info.GroupId = data.GroupId
	info.State = 1
	info.Sys = 0
	return user.GetDB().Create(info).Error
}

// GetByAccount
// 查找用户是否存在
func (user *User) GetByAccount(account string) (*model.UserInfo, error) {
	info := new(model.UserInfo)
	err := user.GetDB().Model(info).Where("account = ?", account).Find(info).Error
	return info, err
}

// GetById
// 根据用户ID查找用户信息
func (user *User) GetById(id string) (*model.UserInfo, error) {
	info := new(model.UserInfo)
	err := user.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// Modify
// 修改用户信息
func (user *User) Modify(info *repository.ModifyUser) error {
	m := make(map[string]any)
	m["account"] = info.Account
	m["name"] = info.Name
	m["password"] = info.Password
	m["mark"] = info.Mark
	m["group_id"] = info.GroupId

	return user.GetDB().Model(&model.UserInfo{}).Where("id = ?", info.ID).Updates(m).Error
}

// Status
// 修改用户信息
func (user *User) Status(info *repository.StatusUser) error {
	return user.GetDB().
		Model(&model.UserInfo{}).
		Where("id = ?", info.ID).
		Update("state", info.State).
		Error
}

// Delete
// 删除用户信息
func (user *User) Delete(id string) error {
	return user.GetDB().Where("id = ?", id).Delete(&model.UserInfo{}).Error
}

// List
// 获取用户列表
func (user *User) List(info *repository.QueryUser) ([]*repository.ResUserGroupData, error) {
	userInfo := new(model.UserInfo)
	list := make([]*repository.ResUserGroupData, 0)

	sqlStr := fmt.Sprintf("SELECT U.*, G.name as group_name FROM %s U, %s G WHERE U.group_id = G.id", userInfo.TableName(), model.UserGroupInfo{}.TableName())
	err := user.DataFilter(fmt.Sprintf("(%s)TB", sqlStr), info, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("id")

		// 用户姓名
		if info.Name != "" {
			query = query.Where("name like ?", fmt.Sprintf("%%%s%%", info.Name))
		}

		// 查询条件
		if info.Condition != "" {
			query = query.
				Where("name like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("mark like ?", fmt.Sprintf("%%%s%%", info.Condition))
		}

		return query, nil
	})
	return list, err
}
