package router

import (
	"go-zone/config"
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
		// 获取文章详情
		v1.GET("/post/:slug", controller.PostItemGetHandler)
		// 获取book列表
		v1.GET("/books", controller.BookListGetHandler)
	}
}

func GetRoute() *gin.Engine {
	return r
}

func init() {
	// 设置gin mode
	env := config.GetConfig().Env
	if env == config.ENV_SLS || env == config.ENV_PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化gin
	r = gin.Default()
	setup()
}
