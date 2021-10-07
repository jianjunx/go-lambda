package main

import (
	"net/http"

	_ "gin-template/boot"
	"gin-template/routers"
)

func main() {
	// 注册路由
	r := routers.Setup()
	http.ListenAndServe(":3000", r)
}
