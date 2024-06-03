package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// PortInfo
// 端口信息
type PortInfo struct {
	Base
	PortBase
}

type PortBase struct {
	Port    int    `gorm:"column:port;type:int;comment:端口号;" json:"port"`                     // 端口号
	State   bool   `gorm:"column:state;type:int;comment:状态 0 停用 1 启用;" json:"state"`          // 状态
	IsTLS   bool   `gorm:"column:is_tls;type:int;comment:是否https;" json:"isTLS"`              // 是否证书加密
	CerId   string `gorm:"column:cert_id;type:nvarchar(100);comment:编码;" json:"certId"`       // 证书编码
	UseGzip bool   `gorm:"column:use_gzip;type:int;comment:使用GZIP 0 停用 1 启用;" json:"useGzip"` // 使用GZIP 0 停用 1 启用
	Mark    string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`           // 备注
}

func (mod *PortInfo) Copy(data *PortBase) {
	mod.Port = data.Port
	mod.IsTLS = data.IsTLS
	mod.CerId = data.CerId
	mod.State = data.State
	mod.UseGzip = data.UseGzip
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (PortInfo) TableName() string {
	return "port_info"
}

func (PortInfo) New() *PortInfo {
	return &PortInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
