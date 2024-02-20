package repository

import "github.com/issueye/grape/internal/common/model"

type CreateGroupMenu struct {
	model.GroupMenuBase
}

type ModifyGroupMenu struct {
	ID string `json:"id"` // 编码
	model.GroupMenuBase
}

// 修改用户组菜单权限
type ModifyGroupMenuState struct {
	Datas []string `json:"datas"` // 需要更新的菜单
}

type StatusGroupMenu struct {
	ID    string `json:"id"`    // 编码
	State uint   `json:"state"` // 备注
}

type QueryGroupMenu struct {
	Condition string `json:"condition" form:"condition"` // 条件
	GroupId   string `json:"groupId" form:"groupId"`     // 用户组编码
	State     uint   `json:"state" form:"state"`         // 备注
	GetAll    bool   `json:"getAll" form:"getAll"`       // 获取所有
	model.Page
}

type ResGroupMenu struct {
	ID       string          `json:"id"`       // 编码
	MenuId   string          `json:"menu_id"`  // 菜单编码
	Name     string          `json:"name"`     // 菜单名称
	Title    string          `json:"title"`    // 菜单标题
	Route    string          `json:"route"`    // 路由
	Icon     string          `json:"icon"`     // 菜单图标
	Auth     int             `json:"auth"`     // 菜单权限级别
	State    uint            `json:"state"`    // 备注
	Children []*ResGroupMenu `json:"children"` // 子菜单
}
