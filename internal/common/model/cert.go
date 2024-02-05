package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// 证书信息
type CertInfo struct {
	Base
	CertBase
}

type CertBase struct {
	Name    string `binding:"required" label:"标题" gorm:"column:name;type:nvarchar(300);comment:名称;" json:"name"` // 名称
	Public  string `label:"公钥" gorm:"column:public;type:nvarchar(300);comment:公有证书路径;" json:"public"`            // 公有证书路径
	Private string `label:"私钥" gorm:"column:private;type:nvarchar(300);comment:私有证书路径;" json:"private"`          // 私有证书路径
	Mark    string `label:"备注" gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                   // 备注
}

func (mod *CertInfo) Copy(data *CertBase) {
	mod.Name = data.Name
	mod.Public = data.Public
	mod.Private = data.Private
	mod.Mark = data.Mark
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
