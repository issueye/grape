package repository

import "github.com/issueye/grape/internal/common/model"

type CreateMenu struct {
	model.MenuBase
}

type ModifyMenu struct {
	ID string `json:"id"` // 编码
	model.MenuBase
}

type StatusMenu struct {
	ID    string `json:"id"`    // 编码
	State uint   `json:"state"` // 备注
}

type QueryMenu struct {
	Condition string `json:"condition" form:"condition"`    // 条件
	Level     int    `json:"level" form:"level,default=-1"` // 等级
	model.Page
}

type ResTreeMenus struct {
	model.Menu
	HasChildren bool          `json:"hasChildren"` // 是否有子集
	Children    []*model.Menu `json:"children"`
}
