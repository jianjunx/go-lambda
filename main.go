package main

import (
	"fmt"
	"net/http"

	_ "go-zone/boot"
	"go-zone/config"
	"go-zone/router"

	"github.com/apex/gateway/v2"
)

func main() {
	// 配置实例
	conf := config.GetConfig()
	// fmt.Println("333", os.Getenv("GOENV_MODE"))
	// 端口号
	port := fmt.Sprintf(":%v", conf.Port)
	// 生产环境初始化lambda
	if conf.Env == config.ENV_SLS {
		// lambda.Init(router.GetRoute())
		gateway.ListenAndServe(port, router.GetRoute())
	} else {
		// 普通服务
		http.ListenAndServe(port, router.GetRoute())
	}
}
