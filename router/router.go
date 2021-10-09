package router

import (
	"go-zone/config"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

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
	// apis
	registerApis()
	// 服务端渲染
	registerHtml()
}
