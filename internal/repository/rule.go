package repository

import "github.com/issueye/grape/internal/common/model"

type CreateRule struct {
	Name      string `json:"name" binding:"required" label:"匹配路由名称"`   // 匹配路由名称
	NodeId    string `json:"nodeId" label:"节点信息编码"`                    // 节点信息编码
	PortId    string `json:"portId" binding:"required" label:"端口信息编码"` // 端口信息编码
	MatchType uint   `json:"matchType" label:"匹配模式"`                   // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Target    string `json:"target" binding:"required" label:"目标地址"`   //  目标服务地址
	Mark      string `json:"mark"`                                     // 备注
}

type ModifyRule struct {
	ID        string `json:"id" binding:"required" label:"编码"`          // 编码
	Name      string `json:"name" binding:"required" label:"匹配路由名称"`    // 匹配路由名称
	NodeId    string `json:"nodeId" label:"节点信息编码"`                     // 节点信息编码
	PortId    string `json:"portId" binding:"required" label:"端口信息编码"`  // 端口信息编码
	MatchType uint   `json:"matchType" binding:"required" label:"匹配模式"` // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Target    string `json:"target" binding:"required" label:"目标地址"`    //  目标服务地址
	Mark      string `json:"mark"`                                      // 备注
}

// 查询信息
type QueryRule struct {
	Conditon string `json:"condition" form:"condition"` // 条件
	NodeId   string `json:"nodeId" form:"nodeId"`       // 节点编码
	PortId   string `json:"portId" form:"portId"`       // 端口编码
	model.Page
}

type QueryRuleRes struct {
	model.RuleInfo
	Port int `gorm:"column:port;type:int;comment:端口号;" json:"port"` // 端口号
}
