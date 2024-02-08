package repository

import "github.com/issueye/grape/internal/common/model"

type CreateGroupMenu struct {
	model.GroupMenuBase
}

type ModifyGroupMenu struct {
	ID string `json:"id"` // 编码
	model.GroupMenuBase
}

type StatusGroupMenu struct {
	ID    string `json:"id"`    // 编码
	State uint   `json:"state"` // 备注
}

type QueryGroupMenu struct {
	Condition string `json:"condition" form:"condition"` // 条件
	GroupId   string `json:"groupId" form:"groupId"`     // 用户组编码
	State     uint   `json:"state" form:"state"`         // 备注
	model.Page
}

type ResGroupMenu struct {
	ID       string          `json:"-"`        // 编码
	Name     string          `json:"name"`     // 菜单名称
	Title    string          `json:"title"`    // 菜单标题
	Route    string          `json:"route"`    // 路由
	Icon     string          `json:"icon"`     // 菜单图标
	Auth     int             `json:"auth"`     // 菜单权限级别
	Children []*ResGroupMenu `json:"children"` // 子菜单
}
