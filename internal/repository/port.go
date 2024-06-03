package repository

import "github.com/issueye/grape/internal/common/model"

// 创建端口信息
type CreatePort struct {
	Port    int    `json:"port" binding:"required" label:"端口号"` // 端口号
	IsTLS   bool   `json:"isTLS"`                               // 是否证书加密
	CertId  string `json:"certId"`                              // 证书编码
	UseGzip bool   `json:"useGzip"`                             // 使用GZIP 0 停用 1 启用
	Mark    string `json:"mark"`                                // 备注
}

// 修改端口信息
type ModifyPort struct {
	ID      string `json:"id" binding:"required" label:"编码"`    // 编码
	Port    int    `json:"port" binding:"required" label:"端口号"` // 端口号
	IsTLS   bool   `json:"isTLS"`                               // 是否证书加密
	CertId  string `json:"certId"`                              // 证书编码
	UseGzip bool   `json:"useGzip"`                             // 使用GZIP 0 停用 1 启用
	Mark    string `json:"mark"`                                // 备注
}

// 查询信息
type QueryPort struct {
	Condition string `json:"condition"` // 条件
	model.Page
}
