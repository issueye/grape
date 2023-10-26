package model

// PortInfo
// 端口信息
type PortInfo struct {
	ID       string `gorm:"column:id;type:nvarchar(100);comment:编码;primaryKey;autoIncrement:false;" json:"id"` // 编码
	Port     int    `gorm:"column:port;type:int;comment:端口号;" json:"port"`                                     // 端口号
	IsTLS    bool   `gorm:"column:is_tls;type:int;comment:是否https;" json:"isTLS"`                              // 是否证书加密
	CertCode string `gorm:"column:cert_code;type:nvarchar(100);comment:编码;" json:"certCode"`                   // 证书编码
	Mark     string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                           // 备注
}

// TableName
// 表名称
func (PortInfo) TableName() string {
	return "port_info"
}

// NodeInfo
// 节点信息
type NodeInfo struct {
	ID       string `gorm:"column:id;type:nvarchar(100);comment:编码;primaryKey;autoIncrement:false;" json:"id"` // 编码
	Name     string `gorm:"column:name;type:nvarchar(300);comment:匹配路由名称;" json:"name"`                        // 匹配路由名称
	NodeType uint   `gorm:"column:node_type;type:int;comment:节点类型 0 api 1 页面;" json:"nodeType"`                // 节点类型 0 api 1 页面
	PagePath string `gorm:"column:page_path;type:int;comment:静态页面存放路径 注：相对路径，由服务对页面进行管理;" json:"pagePath"`     // 静态页面存放路径 注：相对路径，由服务对页面进行管理
	Mark     string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                           // 备注
}

// TableName
// 表名称
func (NodeInfo) TableName() string {
	return "node_info"
}

// RouteInfo
// 路由匹配信息
type RouteInfo struct {
	ID        string `gorm:"column:id;type:nvarchar(100);comment:编码;primaryKey;autoIncrement:false;" json:"id"`            // 编码
	Name      string `gorm:"column:name;type:nvarchar(300);comment:匹配路由名称;" json:"name"`                                   // 匹配路由名称
	MatchType uint   `gorm:"column:match_type;type:int;comment:匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配;" json:"matchType"` // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Target    string `gorm:"column:target;type:nvarchar(2000);comment:目标服务地址;" json:"target"`                              //  目标服务地址
	Mark      string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"`                                      // 备注
}

// TableName
// 表名称
func (RouteInfo) TableName() string {
	return "route_info"
}
