package engine

// http 代理
type HttpProxy struct {
	IsTSL bool // 是否需要证书
	Port  uint // 端口号
}
