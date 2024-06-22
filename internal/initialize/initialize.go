package initialize

import (
	"context"
	"fmt"

	"github.com/issueye/grape/internal/common/nodb"
	"github.com/issueye/grape/internal/config"
	"github.com/issueye/grape/internal/engine"
	"github.com/issueye/grape/internal/global"
)

func Initialize() {
	ctx := context.Background()
	// 初始化运行文件
	InitRuntime()
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
	// 监听报文数据
	nodb.InitDB(ctx)
	// 启动服务
	ShowInfo()
	// 监听服务
	err := global.HttpServer.ListenAndServe()
	if err != nil {
		fmt.Printf("启动服务失败：%v", err)
	}

	// 关闭服务
	global.HttpServer.Shutdown(ctx)
	// 关闭数据库
	close(global.IndexDB)
	// 关闭日志
	global.Logger.Sync()
	// 关闭监听
	ctx.Done()
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
