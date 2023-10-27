package db

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

const LINE_DB_PWD_KEY = "0_1_2_3_4_5_6_7_8_9_))))))))))!!"

type Config struct {
	Username string `json:"user"`     // 用户名称
	Password string `json:"password"` // 密码
	Host     string `json:"host"`     // 服务器地址
	Database string `json:"name"`     // 数据库
	Port     int    `json:"port"`     // 端口号
	LogMode  bool   `json:"logMode"`  // 日志模式
}

// Writer 封装的SQL打印
type Writer struct {
	log    *zap.SugaredLogger
	BPrint bool
}

func (w Writer) Printf(format string, args ...interface{}) {

	if !w.BPrint {
		return
	}

	switch len(args) {
	case 3:
	case 4:
		{
			funcPath := args[0].(string)
			data := strings.Split(funcPath, "/")
			if args[2] == "-" {
				w.log.Debugf("\nSQL语句  %s\n"+
					"执行用时  %0.2fms\n"+
					"影响行数  %s\n"+
					"代码路径  %s\n", args[3], args[1], args[2], data[len(data)-1])
			} else {
				w.log.Debugf("\nSQL语句  %s\n"+
					"执行用时 %0.2fms\n"+
					"影响行数 %d\n"+
					"代码路径 %s\n", args[3], args[1], args[2], data[len(data)-1])
			}
		}
	case 5: // 错误SQL语句
		{
			funcPath := args[0].(string)
			data := strings.Split(funcPath, "/")
			// 判断如果是 SLOW SQL 则使用 warn 级别
			fmt.Println("args[2]", args[2])

			isSlow := false
			switch t := args[2].(type) {
			case int64:
				isSlow = t > 200
			case int:
				isSlow = t > 200
			case float32:
				isSlow = t > 200
			case float64:
				isSlow = t > 200
			default:
				isSlow = false
			}

			if isSlow {
				w.log.Warnf("\nSQL语句  %s\n"+
					"执行情况 SLOW SQL\n"+
					"执行用时 %0.2fms\n"+
					"影响行数 %d\n"+
					"代码路径 %s\n", args[4], args[2], args[3], data[len(data)-1])
			} else {
				w.log.Errorf("\nSQL语句  %s\n"+
					"错误信息  %s\n"+
					"执行用时  %0.2fms\n"+
					"影响行数  %d\n"+
					"代码路径  %s\n", args[4], args[1], args[2], args[3], data[len(data)-1])
			}
		}
	}
}
