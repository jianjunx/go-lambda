package routers

import (
	"gin-template/src/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 注册路由
	v1 := r.Group("/api/v1")
	{
		// 数据
		v1.POST("/webhook/yuque", controllers.PostWebhookYuqueHandler)
		// 数据库数据
		v1.GET("/dynamo", controllers.GetTableNamesHandler)
	}
	return r
}
