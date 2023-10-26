package config

const (
	CfgServerName    = "SERVER-PORT"
	CfgServerMode    = "SERVER-MODE"
	CfgLogMaxSize    = "LOG-MAX-SIZE"
	CfgLogMaxBackups = "LOG-MAX-BACKUPS"
	CfgLogMaxAge     = "LOG-MAX-AGE"
	CfgLogCompress   = "LOG-COMPRESS"
	CfgLogLevel      = "LOG-LEVEL"
)

func InitConfig() {
	SetParamExist(CfgServerName, "10065", "端口号")
	SetParamExist(CfgServerMode, "release", `服务运行模式， debug \ release`)
	SetParamExist(CfgLogMaxSize, "10", "日志大小")
	SetParamExist(CfgLogMaxBackups, "10", "最大备份数")
	SetParamExist(CfgLogMaxAge, "10", "保存天数")
	SetParamExist(CfgLogCompress, "true", "是否压缩")
	SetParamExist(CfgLogLevel, "-1", "日志输出等级")
}
