package repository

import (
	"mime/multipart"

	"github.com/issueye/grape/internal/common/model"
)

type CreateResource struct {
	Title    string `binding:"required" label:"标题" json:"title"` // 标题
	Ext      string `binding:"required" label:"文件类型" json:"ext"` // 文件类型
	FileName string `label:"文件名称" json:"fileName"`               // 文件名称
	Mark     string `label:"备注" json:"mark"`                     // 备注
}

type ModifyResource struct {
	ID       string `json:"id" binding:"required" label:"编码"`    // 编码
	Title    string `binding:"required" label:"标题" json:"title"` // 标题
	Ext      string `binding:"required" label:"文件类型" json:"ext"` // 文件类型
	FileName string `label:"文件名称" json:"fileName"`               // 文件名称
	Mark     string `label:"备注" json:"mark"`                     // 备注
}

// 查询信息
type QueryResource struct {
	Condition string `json:"condition" form:"condition"` // 条件
	model.Page
}

type UploadData struct {
	Type      string                `json:"type" form:"type"`                   // 类型  page 页面文件
	UploadKey *multipart.FileHeader `binding:"required" label:"文件" form:"file"` // 文件
	Id        string                `json:"id" form:"id"`                       // 编码
}
