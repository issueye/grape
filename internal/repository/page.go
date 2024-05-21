package repository

import "github.com/issueye/grape/internal/common/model"

type CreatePage struct {
	Name        string `binding:"required" label:"名称" json:"name"`       // 名称
	Title       string `binding:"required" label:"标题" json:"title"`      // 标题
	PortId      string `binding:"required" label:"端口信息编码" json:"portId"` // 端口信息编码
	ProductCode string `label:"产品代码" json:"productCode"`                 // 产品代码
	Version     string ` label:"版本" json:"version"`                      // 版本
	PagePath    string `json:"pagePath"`                                 // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	Mark        string `json:"mark"`                                     // 备注
}

type ModifyPage struct {
	ID       string `json:"id" binding:"required" label:"编码"`         // 编码
	Name     string `json:"name" binding:"required" label:"名称"`       // 名称
	PortId   string `json:"portId" binding:"required" label:"端口信息编码"` // 端口信息编码
	PagePath string `json:"pagePath"`                                 // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	Mark     string `json:"mark"`                                     // 备注
}

// 查询信息
type QueryPage struct {
	Condition string `json:"condition" form:"condition"` // 条件
	PortId    string `json:"portId" form:"portId"`       // 端口信息编码
	model.Page
}
