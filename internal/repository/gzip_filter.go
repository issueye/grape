package repository

import "github.com/issueye/grape/internal/common/model"

// 创建端口信息
type CreateGzipFilter struct {
	PortId       string `json:"portId"`       // 端口号
	MatchContent string `json:"matchContent"` // 匹配内容
	MatchType    uint   `json:"matchType"`    // 匹配模式
	Mark         string `json:"mark"`         // 备注
}

// 修改端口信息
type ModifyGzipFilter struct {
	ID           string `json:"id" binding:"required" label:"编码"` // 编码
	PortId       string `json:"portId"`                           // 端口号
	MatchContent string `json:"matchContent"`                     // 匹配内容
	MatchType    uint   `json:"matchType"`                        // 匹配模式
	Mark         string `json:"mark"`                             // 备注
}

// 查询信息
type QueryGzipFilter struct {
	Condition string `json:"condition"` // 条件
	PortId    string `json:"portId"`    // 端口号
	model.Page
}
