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
	Name        string `binding:"required" label:"名称" gorm:"column:name;size:300;comment:名称;" json:"name"`                    // 名称
	Title       string `binding:"required" label:"标题" gorm:"column:title;size:300;comment:标题;" json:"title"`                  // 标题
	Version     string `binding:"required" label:"版本" gorm:"column:version;size:50;comment:版本;" json:"version"`               // 版本
	PortId      string `binding:"required" label:"端口号" gorm:"column:port_id;type:int;comment:端口信息编码;" json:"portId"`          // 端口信息编码
	ProductCode string `binding:"required" label:"产品代码" gorm:"column:product_code;size:200;comment:产品代码;" json:"productCode"` // 产品代码
	State       int    `binding:"required" label:"状态" gorm:"column:state;type:int;comment:状态;" json:"state"`                  // 状态
	Mark        string `label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                                      // 备注
}

func (mod *PageInfo) Copy(data *PageBase) {
	mod.Name = data.Name
	mod.Title = data.Title
	mod.Version = data.Version
	mod.PortId = data.PortId
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

type PageVersionInfo struct {
	Base
	PageVersionBase
}

type PageVersionBase struct {
	ProductCode string `binding:"required" label:"产品代码" gorm:"column:product_code;size:200;comment:产品代码;" json:"productCode"` // 产品代码
	Version     string `binding:"required" label:"版本" gorm:"column:version;size:50;comment:版本;" json:"version"`               // 版本
	PagePath    string `label:"页面路径" gorm:"column:page_path;size:2000;comment:页面路径;" json:"pagePath"`                         // 页面路径
	Mark        string `label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                                      // 备注
}

func (mod *PageVersionInfo) Copy(data *PageVersionBase) {
	mod.ProductCode = data.ProductCode
	mod.PagePath = data.PagePath
	mod.Version = data.Version
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (PageVersionInfo) TableName() string {
	return "page_version_info"
}

func (PageVersionInfo) New() *PageVersionInfo {
	return &PageVersionInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
