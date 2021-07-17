package main

import (
	"fmt"

	"gin-template/logger"
	"gin-template/settings"
)

func main() {
	// 读取配置
	err := settings.Init()
	if err != nil {
		fmt.Printf("配置初始化失败 %v\n", err)
		return
	}
	// 初始化日志
	err = logger.Init(&settings.Config)
	if err != nil {
		fmt.Printf("Logger初始化失败 %v\n", err)
		return
	}
	// 初始化MySQL链接
	// 初始化Redis链接
	// 注册路由
	// 启动服务（优雅关机）
}
