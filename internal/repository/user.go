package repository

type CreateUser struct {
	Account  string `json:"account"`  // uid 登录名
	Name     string `json:"name"`     // 用户姓名
	Password string `json:"password"` // 密码
	Mark     string `json:"mark"`     // 备注
}

type ModifyUser struct {
	ID       int64  `json:"id"`       // 编码
	Account  string `json:"account"`  // uid 登录名
	Name     string `json:"name"`     // 用户姓名
	Password string `json:"password"` // 密码
	Mark     string `json:"mark"`     // 备注
}

type StatusUser struct {
	ID    int64 `json:"id"`    // 编码
	State uint  `json:"state"` // 备注
}

type QueryUser struct {
	Account string `json:"account"` // uid 登录名
	Name    string `json:"name"`    // 用户姓名
	Mark    string `json:"mark"`    // 备注
}

// Login
// 用户登录
type Login struct {
	Account  string `json:"account"`  // 登录名
	Password string `json:"password"` // 密码
}
