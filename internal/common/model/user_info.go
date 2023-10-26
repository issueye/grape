package model

// User
// 用户信息
type User struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:false;type:int" json:"id"` // 编码
	Account  string `gorm:"column:account;type:nvarchar(50)" json:"account"`             // uid 登录名
	Name     string `gorm:"column:name;type:nvarchar(50)" json:"name"`                   // 用户姓名
	Password string `gorm:"column:password;type:nvarchar(50)" json:"password"`           // 密码
	Mark     string `gorm:"column:mark;type:nvarchar(500)" json:"mark"`                  // 备注
	State    uint   `gorm:"column:state;type:int" json:"state"`                          // 状态 0 停用 1 启用
}

func (User) TableName() string {
	return "user_info"
}
