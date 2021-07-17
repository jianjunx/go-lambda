package routers

import (
	"gin-template/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 注册中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册路由
	r.GET("/", func(c *gin.Context) {})
	return r
}
