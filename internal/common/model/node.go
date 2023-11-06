package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// NodeInfo
// 节点信息
type NodeInfo struct {
	Base
	Name     string `gorm:"column:name;type:nvarchar(300);comment:匹配路由名称;" json:"name"`                               // 匹配路由名称
	PortId   string `gorm:"column:port_id;type:nvarchar(100);comment:端口信息编码;" json:"portId"`                          // 端口信息编码
	NodeType uint   `gorm:"column:node_type;type:int;comment:节点类型 0 api 1 页面;" json:"nodeType"`                       // 节点类型 0 api 1 页面
	Target   string `gorm:"column:target;type:nvarchar(2000);comment:目标服务地址;" json:"target"`                          //  目标服务地址
	PagePath string `gorm:"column:page_path;type:nvarchar(2000);comment:静态页面存放路径 注：相对路径，由服务对页面进行管理;" json:"pagePath"` // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	FileName string `gorm:"column:file_name;type:nvarchar(2000);comment:文件名称;" json:"fileName"`                       // 文件名称
	Mark     string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                                  // 备注
}

// TableName
// 表名称
func (NodeInfo) TableName() string {
	return "node_info"
}

func (NodeInfo) New() *NodeInfo {
	return &NodeInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
