package repository

// 创建端口信息
type CreatePortInfo struct {
	ID       string `json:"id"`       // 编码
	Port     int    `json:"port"`     // 端口号
	IsTLS    bool   `json:"isTLS"`    // 是否证书加密
	CertCode string `json:"certCode"` // 证书编码
	Mark     string `json:"mark"`     // 备注
}
