package global

const SYS_AUTO_CREATE = "系统自定生成"

const TokenHeadName = "Bearer" // Token 认证方式

const AdminGroupId string = "10000"
const AdminGroupName string = "admin"
const AdminId string = "10000"
const AdminName string = "admin"
const AdminAccount string = "admin"

type RouterGroupName string

const (
	RGN_user       RouterGroupName = "user"
	RGN_userGroup  RouterGroupName = "userGroup"
	RGN_menu       RouterGroupName = "menu"
	RGN_menuGroup  RouterGroupName = "groupMenu"
	RGN_target     RouterGroupName = "target"
	RGN_rule       RouterGroupName = "rule"
	RGN_port       RouterGroupName = "port"
	RGN_page       RouterGroupName = "page"
	RGN_resource   RouterGroupName = "resource"
	RGN_gzipFilter RouterGroupName = "gzipFilter"
)
