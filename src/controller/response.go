package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 响应结果错误
func respError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": 1,
		"msg": err.Error(),
	})
	zap.L().Error("ERR", zap.Error(err))
	c.Abort()
}

// 响应结果成功
func respSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"err":  0,
		"data": data,
		"msg":  "ok",
	})
}
