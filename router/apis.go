package router

import "go-zone/src/controller"

func registerApis() {
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
