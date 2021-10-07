package router

import (
	"go-zone/src/controller"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func setup() {
	// 注册路由
	v1 := r.Group("/api/v1")
	{
		// 数据
		v1.POST("/webhook/yuque", controller.PostWebhookYuqueHandler)
		// 获取列表数据
		v1.GET("/posts", controller.PostListGetHandler)
	}
}

func GetRoute() *gin.Engine {
	return r
}

func init() {
	r = gin.Default()
	setup()
}
