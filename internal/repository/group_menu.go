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
	model.Page
}
