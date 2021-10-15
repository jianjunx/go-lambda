package controller

import (
	"go-zone/packed/utils"
	"go-zone/src/model"
	"go-zone/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HtmlHomeHandler(c *gin.Context) {
	books, err := service.BookListGet()
	if err != nil {
		respError(c, err)
		return
	}
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
	//render with master
	c.HTML(http.StatusOK, "index", gin.H{
		"title":  "Index title!",
		"books":  books,
		"posts":  list,
		"total":  total,
		"params": params,
		"pagers": utils.GetPages(float64(total), float64(params.Size)),
	})
}

func HtmlPostDetailHandler(c *gin.Context) {
	books, err := service.BookListGet()
	if err != nil {
		respError(c, err)
		return
	}
	slue := c.Param("slue")
	detail, err := service.PostItemGet(slue)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "page.tpl", gin.H{
		"title":  "Index title!",
		"books":  books,
		"detail": detail,
	})
}
