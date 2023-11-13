package initialize

import (
	"fmt"

	"github.com/issueye/grape/internal/config"
	"github.com/issueye/grape/internal/engine"
	"github.com/issueye/grape/internal/global"
	"github.com/panjf2000/ants/v2"
)

func Initialize() {
	// 初始化运行文件
	InitRuntime()

	// 初始化
	// 创建一个协程池
	pool, err := ants.NewPool(2048) // 这里的 10 表示协程池中的初始协程数量
	if err != nil {
		panic(err)
	}

	global.GoAnts = pool

	// 配置参数
	config.InitConfig()
	// 日志
	InitLog()
	// 数据
	InitData()
	// 启动引擎
	engine.Start()
	// http服务
	InitServer()
	// 启动服务
	ShowInfo()
	// 监听服务
	_ = global.HttpServer.ListenAndServe()
}

var (
	AppName string
	Branch  string
	Commit  string
	Date    string
	Version string
)

func ShowInfo() {
	bannerStr := `
	▄███▄██   ██▄████   ▄█████▄  ██▄███▄    ▄████▄  
	██▀  ▀██   ██▀       ▀ ▄▄▄██  ██▀  ▀██  ██▄▄▄▄██ 
	██    ██   ██       ▄██▀▀▀██  ██    ██  ██▀▀▀▀▀▀ 
	▀██▄▄███   ██       ██▄▄▄███  ███▄▄██▀  ▀██▄▄▄▄█ 
	 ▄▀▀▀ ██   ▀▀        ▀▀▀▀ ▀▀  ██ ▀▀▀      ▀▀▀▀▀  
	 ▀████▀▀                      ██                 
   
	代理管理服务 grape
	`
	fmt.Println(bannerStr) // mona12 风格

	info := `
	AppName: %s
	Branch : %s
	Commit : %s
	Date   : %s
	Version: %s
	
	`
	fmt.Printf(info+"\n", AppName, Branch, Commit, Date, Version)
}
