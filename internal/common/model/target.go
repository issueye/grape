package model

import (
	"strconv"

	"github.com/issueye/grape/pkg/utils"
)

// 目标地址信息
type TargetInfo struct {
	Base
	TargetBase
}

type TargetBase struct {
	Name string `gorm:"column:name;type:nvarchar(300);comment:目标地址;" json:"name"` // 目标地址
	Mark string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`  // 备注
}

func (mod *TargetInfo) Copy(data *TargetBase) {
	mod.Name = data.Name
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (TargetInfo) TableName() string {
	return "target_info"
}

func (TargetInfo) New() *TargetInfo {
	return &TargetInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
