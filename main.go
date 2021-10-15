package main

import (
	"fmt"
	"net/http"

	_ "go-zone/boot"
	"go-zone/config"
	"go-zone/router"

	"github.com/apex/gateway/v2"
	"go.uber.org/zap"
)

func main() {
	// 配置实例
	conf := config.GetConfig()

	// 端口号
	port := fmt.Sprintf(":%v", conf.Port)
	// 生产环境初始化lambda
	if conf.Env == config.ENV_SLS {
		// lambda.Init(router.GetRoute())
		gateway.ListenAndServe(port, router.GetRoute())
		return
	}

	// log
	lg := zap.L()
	defer lg.Sync()

	// 普通服务
	http.ListenAndServe(port, router.GetRoute())
}
