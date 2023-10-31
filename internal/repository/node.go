package repository

import "github.com/issueye/grape/internal/common/model"

type CreateNode struct {
	Name     string `json:"name" binding:"required" label:"节点名称"`      // 匹配路由名称
	PortId   string `json:"portId" binding:"required" label:"端口信息编码"`  // 端口信息编码
	Target   string `json:"target"  binding:"required" label:"目标服务地址"` //  目标服务地址
	NodeType uint   `json:"nodeType" label:"节点类型"`                     // 节点类型 0 api 1 页面
	PagePath string `json:"pagePath"`                                  // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	Mark     string `json:"mark"`                                      // 备注
}

type ModifyNode struct {
	ID       string `json:"id" binding:"required" label:"编码"`          // 编码
	Name     string `json:"name" binding:"required" label:"节点名称"`      // 匹配路由名称
	PortId   string `json:"portId" binding:"required" label:"端口信息编码"`  // 端口信息编码
	Target   string `json:"target"  binding:"required" label:"目标服务地址"` //  目标服务地址
	NodeType uint   `json:"nodeType" label:"节点类型"`                     // 节点类型 0 api 1 页面
	PagePath string `json:"pagePath"`                                  // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	Mark     string `json:"mark"`                                      // 备注
}

// 查询信息
type QueryNode struct {
	Conditon string `json:"condition" form:"condition"` // 条件
	PortId   string `json:"portId" form:"portId"`       // 端口信息编码
	model.Page
}
