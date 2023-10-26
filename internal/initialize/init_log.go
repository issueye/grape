package initialize

import (
	"path/filepath"

	"github.com/issueye/grape/internal/config"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/pkg/logger"
)

func InitLog() {
	logConf := new(logger.Config)
	logConf.Path = filepath.Join("runtime", "logs")
	logConf.MaxSize = config.GetParam(config.CfgLogMaxSize, "10").Int()
	logConf.MaxBackups = config.GetParam(config.CfgLogMaxBackups, "10").Int()
	logConf.MaxAge = config.GetParam(config.CfgLogMaxAge, "10").Int()
	logConf.Compress = config.GetParam(config.CfgLogCompress, "true").Bool()
	logConf.Level = config.GetParam(config.CfgLogLevel, "-1").Int()
	global.Log, global.Logger = logger.InitLogger(logConf)
}
