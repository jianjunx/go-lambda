package main

import (
	"fmt"
	"net/http"

	_ "go-zone/boot"
	"go-zone/config"
	"go-zone/packed/lambda"
	"go-zone/router"
)

func main() {
	// 配置实例
	conf := config.GetConfig()
	// 生产环境初始化lambda
	if conf.Env == config.ENV_PROD {
		lambda.Init(router.GetRoute())
		return
	}
	// 普通服务
	port := fmt.Sprintf(":%v", conf.Port)
	http.ListenAndServe(port, router.GetRoute())
}
