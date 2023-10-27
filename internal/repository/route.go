package repository

import "github.com/issueye/grape/internal/common/model"

type CreateRoute struct {
	Name      string `json:"name" binding:"required" label:"匹配路由名称"`    // 匹配路由名称
	NodeId    string `json:"nodeId" binding:"required" label:"节点信息编码"`  // 节点信息编码
	MatchType uint   `json:"matchType" binding:"required" label:"匹配模式"` // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Target    string `json:"target" binding:"required" label:"目标地址"`    //  目标服务地址
	Mark      string `json:"mark"`                                      // 备注
}

type ModifyRoute struct {
	ID        string `json:"id" binding:"required" label:"编码"`          // 编码
	Name      string `json:"name" binding:"required" label:"匹配路由名称"`    // 匹配路由名称
	NodeId    string `json:"nodeId" binding:"required" label:"节点信息编码"`  // 节点信息编码
	MatchType uint   `json:"matchType" binding:"required" label:"匹配模式"` // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Target    string `json:"target" binding:"required" label:"目标地址"`    //  目标服务地址
	Mark      string `json:"mark"`                                      // 备注
}

// 查询信息
type QueryRoute struct {
	Conditon string `json:"condition"` // 条件
	NodeId   string `json:"nodeId"`    // 节点编码
	model.Page
}
