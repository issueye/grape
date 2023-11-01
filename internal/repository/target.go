package repository

import "github.com/issueye/grape/internal/common/model"

type CreateTarget struct {
	Name string `json:"name" binding:"required" label:"目标地址"` // 目标地址
	Mark string `json:"mark"`                                 // 备注
}

type ModifyTarget struct {
	ID   string `json:"id" binding:"required" label:"编码"`     // 编码
	Name string `json:"name" binding:"required" label:"目标地址"` // 目标地址
	Mark string `json:"mark"`                                 // 备注
}

// 查询信息
type QueryTarget struct {
	Condition string `json:"condition" form:"condition"` // 条件
	model.Page
}
