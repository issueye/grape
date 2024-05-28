package repository

import "github.com/issueye/grape/internal/common/model"

type CreatePage struct {
	Name            string `binding:"required" label:"名称" json:"name"`       // 名称
	Title           string `binding:"required" label:"标题" json:"title"`      // 标题
	PortId          string `binding:"required" label:"端口信息编码" json:"portId"` // 端口信息编码
	ProductCode     string `label:"产品代码" json:"productCode"`                 // 产品代码
	Version         string `label:"版本" json:"version"`                      // 版本
	UseVersionRoute int    `label:"使用版本路由" json:"useVersionRoute"`           // 使用版本路由
	PagePath        string `json:"pagePath"`                                 // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	Mark            string `json:"mark"`                                     // 备注
}

type ModifyPage struct {
	ID              string `binding:"required" label:"编码" json:"id"`         // 编码
	Name            string `binding:"required" label:"名称" json:"name"`       // 名称
	PortId          string `binding:"required" label:"端口信息编码" json:"portId"` // 端口信息编码
	PagePath        string `label:"" json:"pagePath"`                        // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	Version         string `label:"版本" json:"version"`                       // 版本
	UseVersionRoute int    `label:"使用版本路由" json:"useVersionRoute"`           // 使用版本路由
	Thumbnail       string `label:"缩略图" json:"thumbnail"`                    // 缩略图
	Mark            string `label:"备注" json:"mark"`                          // 备注
}

// 查询信息
type QueryPage struct {
	Condition string `json:"condition" form:"condition"` // 条件
	PortId    string `json:"portId" form:"portId"`       // 端口信息编码
	model.Page
}

// 查询信息
type QueryPageVersion struct {
	Condition   string `json:"condition" form:"condition"`     // 条件
	ProductCode string `json:"productCode" form:"productCode"` // 产品编码
	PortId      string `json:"portId" form:"portId"`           // 端口信息编码
	model.Page
}
