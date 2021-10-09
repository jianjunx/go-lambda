package controller

import (
	"fmt"
	"go-zone/packed/utils"
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
