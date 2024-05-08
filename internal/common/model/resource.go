package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// ResourceInfo
// 资源信息
type ResourceInfo struct {
	Base
	ResourceBase
}

type ResourceBase struct {
	Name     string `binding:"required" label:"名称" gorm:"column:name;size:300;comment:名称;" json:"name"`       // 名称
	Title    string `binding:"required" label:"标题" gorm:"column:title;size:300;comment:标题;" json:"title"`     // 标题
	Folder   string `binding:"required" label:"文件夹" gorm:"column:folder;size:300;comment:文件夹;" json:"folder"` // 文件夹
	FileName string `label:"文件名称" gorm:"column:file_name;size:2000;comment:文件名称;" json:"fileName"`            // 文件名称
	Mark     string `label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                         // 备注
}

func (mod *ResourceInfo) Copy(data *ResourceBase) {
	mod.Name = data.Name
	mod.Title = data.Title
	mod.Folder = data.Folder
	mod.FileName = data.FileName
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (ResourceInfo) TableName() string {
	return "resource_info"
}

func (ResourceInfo) New() *ResourceInfo {
	return &ResourceInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
