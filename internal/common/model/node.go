package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// NodeInfo
// 节点信息
type NodeInfo struct {
	Base
	NodeBase
}

type NodeBase struct {
	Name     string `binding:"required" label:"节点名称" gorm:"column:name;type:nvarchar(300);comment:匹配路由名称;" json:"name"`              // 匹配路由名称
	PortId   string `binding:"required" label:"端口号" gorm:"column:port_id;type:nvarchar(100);comment:端口信息编码;" json:"portId"`          // 端口信息编码
	NodeType uint   `label:"节点类型" gorm:"column:node_type;type:int;comment:节点类型 0 api 1 页面;" json:"nodeType"`                         // 节点类型 0 api 1 页面
	Target   string `label:"目标服务地址" gorm:"column:target;type:nvarchar(2000);comment:目标服务地址;" json:"target"`                          //  目标服务地址
	PagePath string `label:"静态页面路径" gorm:"column:page_path;type:nvarchar(2000);comment:静态页面存放路径 注：相对路径，由服务对页面进行管理;" json:"pagePath"` // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	FileName string `label:"文件名称" gorm:"column:file_name;type:nvarchar(2000);comment:文件名称;" json:"fileName"`                         // 文件名称
	Mark     string `label:"备注" gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                                      // 备注
}

func (mod *NodeInfo) Copy(data *NodeBase) {
	mod.Name = data.Name
	mod.PortId = data.PortId
	mod.Target = data.Target
	mod.PagePath = data.PagePath
	mod.FileName = data.FileName
	mod.Mark = data.Mark
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
