package repository

import "github.com/issueye/grape/internal/common/model"

type CreateResource struct {
	Name     string `binding:"required" label:"名称" json:"name"`    // 名称
	Title    string `binding:"required" label:"标题" json:"title"`   // 标题
	Folder   string `binding:"required" label:"文件夹" json:"folder"` // 文件夹
	FileName string `label:"文件名称" json:"fileName"`                 // 文件名称
	Mark     string `label:"备注" json:"mark"`                       // 备注
}

type ModifyResource struct {
	ID       string `json:"id" binding:"required" label:"编码"`      // 编码
	Name     string `binding:"required" label:"名称" json:"name"`    // 名称
	Title    string `binding:"required" label:"标题" json:"title"`   // 标题
	Folder   string `binding:"required" label:"文件夹" json:"folder"` // 文件夹
	FileName string `label:"文件名称" json:"fileName"`                 // 文件名称
	Mark     string `label:"备注" json:"mark"`                       // 备注
}

// 查询信息
type QueryResource struct {
	Condition string `json:"condition" form:"condition"` // 条件
	model.Page
}
