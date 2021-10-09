package controller

import (
	"errors"
	"fmt"
	"go-zone/packed/utils"
	"go-zone/src/model"
	"go-zone/src/service"
	"os"

	"github.com/gin-gonic/gin"
)

// 语雀webhook post请求的body
type addPostBody struct {
	Data model.PostAttributes `json:"data"`
}

// 处理语雀webhook请求
func PostWebhookYuqueHandler(c *gin.Context) {
	secret := c.Query("secret")
	envSecret := os.Getenv("WEBHOOK_SECRET")
	if secret != envSecret {
		respError(c, errors.New("secret 验证失败"))
		return
	}
	p := addPostBody{}
	err := c.BindJSON(&p)
	if err != nil {
		respError(c, err)
		return
	}
	err = service.PostItemDataAdd(&p.Data)
	if err != nil {
		respError(c, err)
		return
	}
	respSuccess(c, nil)
}

// 文章列表
func PostListGetHandler(c *gin.Context) {
	params := model.PostSearchParam{
		Page: utils.StrToInt(c.Query("page"), 1),
		Size: utils.StrToInt(c.Query("size"), 10),
		Book: utils.StrToInt(c.Query("book"), 0),
	}
	list, total, err := service.PostListGet(&params)
	if err != nil {
		respError(c, err)
		return
	}
	respSuccess(c, &model.PageListStruct{
		Size:  params.Size,
		Page:  params.Page,
		Total: total,
		List:  list,
	})
}

// 文章详情
func PostItemGetHandler(c *gin.Context) {
	fmt.Println("req", c.Request.Host)
	slug := c.Param("slug")
	if slug == "" {
		respError(c, errors.New("缺少参数"))
		return
	}
	post, err := service.PostItemGet(slug)
	if err != nil {
		respError(c, err)
		return
	}
	respSuccess(c, post)
}