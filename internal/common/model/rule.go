package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// RouteInfo
// 路由匹配信息
type RuleInfo struct {
	Base
	Name        string `gorm:"column:name;type:nvarchar(300);comment:匹配路由名称;" json:"name"`                                   // 匹配路由名称
	TargetId    string `gorm:"column:target_id;type:nvarchar(300);comment:目标路由;" json:"targetId"`                            // 目标地址编码
	TargetRoute string `gorm:"column:target_route;type:nvarchar(300);comment:目标路由;" json:"targetRoute"`                      // 目标路由
	PortId      string `gorm:"column:port_id;type:nvarchar(100);comment:端口信息编码;" json:"portId"`                              // 端口信息编码
	Method      string `gorm:"column:method;type:nvarchar(100);comment:请求方法;" json:"method"`                                 // 请求方法
	NodeId      string `gorm:"column:node_id;type:nvarchar(100);comment:节点信息编码;" json:"nodeId"`                              // 节点编码
	MatchType   uint   `gorm:"column:match_type;type:int;comment:匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配;" json:"matchType"` // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Mark        string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                                      // 备注
}

// TableName
// 表名称
func (RuleInfo) TableName() string {
	return "rule_info"
}

func (RuleInfo) New() *RuleInfo {
	return &RuleInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
