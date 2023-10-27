package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// RouteInfo
// 路由匹配信息
type RouteInfo struct {
	Base
	Name      string `gorm:"column:name;type:nvarchar(300);comment:匹配路由名称;" json:"name"`                                   // 匹配路由名称
	NodeId    string `gorm:"column:node_id;type:nvarchar(100);comment:端口信息编码;" json:"nodeId"`                              // 节点信息
	MatchType uint   `gorm:"column:match_type;type:int;comment:匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配;" json:"matchType"` // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Mark      string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                                      // 备注
}

// TableName
// 表名称
func (RouteInfo) TableName() string {
	return "route_info"
}

func (RouteInfo) New() *RouteInfo {
	return &RouteInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
