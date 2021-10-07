package controllers

import (
	"gin-template/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTableNamesHandler(c *gin.Context) {
	names, err := service.PrintDynamoName()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": names,
	})
}
