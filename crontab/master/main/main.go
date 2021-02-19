package main

import (
	"crontab/master"
	"flag"
	"fmt"
	"runtime"
	"time"
)

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置线程和核心数相同
}

var (
	confFile string // 配置文件路径
)

func initArgs() {
	// master -config ./master.json
	flag.StringVar(&confFile, "config", "./master/main/master.json", "指定master.json配置文件")
	flag.Parse() // 解析并赋值
}

func main() {
	var (
		err error
	)

	// 初始化线程
	initEnv()

	// 解析命令行参数
	initArgs()

	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	// 服务发现
	if err = master.InitWorkerMgr(); err != nil {
		goto ERR
	}

	// 日志管理器
	if err = master.InitLogMgr(); err !=nil {
		goto ERR
	}

	// 任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	// 启动Api Http服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)

}
