package repository

import "github.com/issueye/grape/internal/common/model"

type CreateUserGroup struct {
	Name string `json:"name"` // 组名称
	Mark string `json:"mark"` // 备注
}

type ModifyUserGroup struct {
	ID   string `json:"id"`   // 编码
	Name string `json:"name"` // 组名称
	Mark string `json:"mark"` // 备注
}

type StatusUserGroup struct {
	ID    string `json:"id"`    // 编码
	State uint   `json:"state"` // 备注
}

type QueryUserGroup struct {
	Name string `json:"name"` // 组名称
	Mark string `json:"mark"` // 备注
	model.Page
}
