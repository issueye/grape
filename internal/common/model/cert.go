package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// 证书信息
type CertInfo struct {
	Base
	Name    string `gorm:"column:name;type:nvarchar(300);comment:名称;" json:"name"`           // 名称
	Public  string `gorm:"column:public;type:nvarchar(300);comment:公有证书路径;" json:"public"`   // 公有证书路径
	Private string `gorm:"column:private;type:nvarchar(300);comment:私有证书路径;" json:"private"` // 私有证书路径
	Mark    string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`          // 备注
}

// TableName
// 表名称
func (CertInfo) TableName() string {
	return "cert_info"
}

func (CertInfo) New() *CertInfo {
	return &CertInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
