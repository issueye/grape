package repository

import "github.com/issueye/grape/internal/common/model"

type CreateUser struct {
	Account  string `json:"account"`  // uid 登录名
	Name     string `json:"name"`     // 用户姓名
	Password string `json:"password"` // 密码
	GroupId  string `json:"groupId"`  // 用户组
	Mark     string `json:"mark"`     // 备注
}

type ModifyUser struct {
	ID       string `json:"id"`       // 编码
	Account  string `json:"account"`  // uid 登录名
	Name     string `json:"name"`     // 用户姓名
	Password string `json:"password"` // 密码
	GroupId  string `json:"groupId"`  // 用户组
	Mark     string `json:"mark"`     // 备注
}

type StatusUser struct {
	ID    string `json:"id"`    // 编码
	State uint   `json:"state"` // 备注
}

type QueryUser struct {
	Conditon string `json:"condition" form:"condition"` // 条件
	Account  string `json:"account" form:"account"`     // uid 登录名
	Name     string `json:"name" form:"name"`           // 用户姓名
	Mark     string `json:"mark" form:"mark"`           // 备注
	model.Page
}

type ResUserGroupData struct {
	model.UserInfo
	Name string `gorm:"column:group_name" json:"groupName"` // 组名称
}

// Login
// 用户登录
type Login struct {
	Account  string `json:"account"`  // 登录名
	Password string `json:"password"` // 密码
}
