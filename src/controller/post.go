package controller

import (
	"fmt"
	"go-zone/src/model"
	"go-zone/src/service"

	"github.com/gin-gonic/gin"
)

// 语雀webhook post请求的body
type addPostBody struct {
	Data model.PostAttributes `json:"data"`
}

// 处理语雀webhook请求
func PostWebhookYuqueHandler(c *gin.Context) {
	p := addPostBody{}
	err := c.BindJSON(&p)
	if err != nil {
		respError(c, err)
		return
	}
	err = service.PostItemDataAdd(&p.Data)
	if err != nil {
		fmt.Println("add data err: ", err)
		respError(c, err)
		return
	}
	fmt.Println("success")
	respSuccess(c, nil)
}

// post列表
func PostListGetHandler(c *gin.Context) {
	list, err := service.PostListGet()
	if err != nil {
		respError(c, err)
		return
	}
	respSuccess(c, list)
}
