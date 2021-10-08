package main

import (
	"fmt"
	"net/http"
	"os"

	_ "go-zone/boot"
	"go-zone/config"
	"go-zone/packed/lambda"
	"go-zone/router"
)

func main() {
	// 配置实例
	conf := config.GetConfig()
	fmt.Println("333", os.Getenv("GOENV_MODE"))
	// 生产环境初始化lambda
	if conf.Env == config.ENV_SLS {
		lambda.Init(router.GetRoute())
		return
	}
	// 普通服务
	port := fmt.Sprintf(":%v", conf.Port)
	http.ListenAndServe(port, router.GetRoute())
}
