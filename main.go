package main

import (
	"fmt"
	"net/http"

	_ "go-zone/boot"
	"go-zone/config"
	"go-zone/router"
)

func main() {
	// 注册路由
	r := router.Setup()
	http.ListenAndServe(fmt.Sprintf(":%v", config.GetConfig().Port), r)
}
