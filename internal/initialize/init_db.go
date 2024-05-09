package initialize

import (
	"fmt"
	"path/filepath"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/logic"
	"github.com/issueye/grape/internal/service"
	"github.com/issueye/grape/pkg/db"
)

// 初始化其他数据
func InitData() {
	path := filepath.Join("runtime", "data", "data.db")
	global.DB = db.InitSqlite(path, global.Log)

	// 初始化表
	err := global.DB.AutoMigrate(
		&model.UserInfo{},      // 用户
		&model.UserGroupInfo{}, // 用户组
		&model.Menu{},          // 菜单
		&model.GroupMenu{},     // 用户组菜单权限
		&model.PortInfo{},      // 端口信息
		&model.PageInfo{},      // 页面信息
		&model.RuleInfo{},      // 规则信息
		&model.CertInfo{},      // 证书信息
		&model.TargetInfo{},    // 目标服务地址信息
		&model.ResourceInfo{},  // 资源信息
	)

	if err != nil {
		panic(fmt.Errorf("初始化表失败 %s", err.Error()))
	}

	err = service.NewUserGroup().CreateAdminNonExistent()
	if err != nil {
		panic("检查管理员用户组信息失败 " + err.Error())
	}

	// 创建 admin 用户
	err = service.NewUser().CreateAdminNonExistent()
	if err != nil {
		panic("初始化数据失败 " + err.Error())
	}

	// 菜单元数据
	logic.Menu{}.CreateMenuNonExistent()

	// 权限组菜单数据
	logic.UserGroup{}.CreateAdminNonExistent()
}
