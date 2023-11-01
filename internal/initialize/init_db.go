package initialize

import (
	"fmt"
	"path/filepath"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/service"
	"github.com/issueye/grape/pkg/db"
)

// 初始化其他数据
func InitData() {
	path := filepath.Join("runtime", "data", "data.db")
	global.DB = db.InitSqlite(path, global.Log)

	// 初始化表
	err := global.DB.AutoMigrate(
		&model.User{},
		&model.PortInfo{},
		&model.NodeInfo{},
		&model.RuleInfo{},
		&model.CertInfo{},
		&model.TargetInfo{},
	)

	if err != nil {
		panic(fmt.Errorf("初始化表失败 %s", err.Error()))
	}

	// 创建 admin 用户
	err = service.NewUser(global.DB).CreateAdminNonExistent()
	if err != nil {
		panic("初始化数据失败 " + err.Error())
	}
}
