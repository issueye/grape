package repository

import "github.com/issueye/grape/internal/common/model"

type CreateRule struct {
	Name        string `json:"name" binding:"required" label:"匹配路由名称"`     // 匹配路由名称
	PortId      int    `json:"portId" binding:"required" label:"端口信息编码"`   // 端口信息编码
	Method      string `json:"method" binding:"required" label:"请求方法"`     // 请求方法
	MatchType   uint   `json:"matchType" label:"匹配模式"`                     // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	TargetId    string `json:"targetId" binding:"required" label:"目标地址编码"` //  目标服务地址
	TargetRoute string `json:"targetRoute" label:"目标路由"`                   // 目标路由
	Mark        string `json:"mark"`                                       // 备注
}

type ModifyRule struct {
	ID          string `json:"id" binding:"required" label:"编码"`           // 编码
	Name        string `json:"name" binding:"required" label:"匹配路由名称"`     // 匹配路由名称
	PortId      int    `json:"portId" binding:"required" label:"端口信息编码"`   // 端口信息编码
	Method      string `json:"method" binding:"required" label:"请求方法"`     // 请求方法
	MatchType   uint   `json:"matchType" label:"匹配模式"`                     // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	TargetId    string `json:"targetId" binding:"required" label:"目标地址编码"` //  目标服务地址
	TargetRoute string `json:"targetRoute" label:"目标路由"`                   // 目标路由
	Mark        string `json:"mark"`                                       // 备注
}

// 查询信息
type QueryRule struct {
	Condition string `json:"condition" form:"condition"` // 条件
	PortId    string `json:"portId" form:"portId"`       // 端口编码
	MatchType uint   `json:"matchType" form:"matchType"` // 匹配类型 0 所有 1 gin 2 mux
	model.Page
}

type QueryRuleRes struct {
	model.RuleInfo
	Port   int    `gorm:"column:port;type:int;comment:端口号;" json:"port"`                // 端口号
	Target string `gorm:"column:target;type:nvarchar(300);comment:目标地址;" json:"target"` // 目标地址
	Node   string `gorm:"column:node;type:nvarchar(100);comment:节点信息;" json:"node"`     // 节点
}
