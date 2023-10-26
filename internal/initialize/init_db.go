package initialize

import (
	"path/filepath"

	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/service"
	"github.com/issueye/grape/pkg/db"
)

// 初始化其他数据
func InitData() {
	path := filepath.Join("runtime", "data", "data.db")
	global.DB = db.InitSqlite(path, global.Log)

	err := service.NewUser(global.DB).CreateAdminNonExistent()
	if err != nil {
		panic("初始化数据失败，失败原因：" + err.Error())
	}
}
