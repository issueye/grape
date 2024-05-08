package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// PageInfo
// 页面信息
type PageInfo struct {
	Base
	PageBase
}

type PageBase struct {
	Name     string `binding:"required" label:"名称" gorm:"column:name;size:300;comment:名称;" json:"name"`                     // 名称
	Title    string `binding:"required" label:"标题" gorm:"column:title;size:300;comment:标题;" json:"title"`                   // 标题
	Version  string `binding:"required" label:"版本" gorm:"column:version;size:50;comment:版本;" json:"version"`                // 版本
	PortId   string `binding:"required" label:"端口号" gorm:"column:port_id;type:nvarchar(100);comment:端口信息编码;" json:"portId"` // 端口信息编码
	FileName string `label:"文件名称" gorm:"column:file_name;type:nvarchar(2000);comment:文件名称;" json:"fileName"`                // 文件名称
	Mark     string `label:"备注" gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                             // 备注
}

func (mod *PageInfo) Copy(data *PageBase) {
	mod.Name = data.Name
	mod.Title = data.Title
	mod.Version = data.Version
	mod.PortId = data.PortId
	mod.FileName = data.FileName
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (PageInfo) TableName() string {
	return "page_info"
}

func (PageInfo) New() *PageInfo {
	return &PageInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
