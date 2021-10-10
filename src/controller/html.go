package controller

import (
	"fmt"
	"go-zone/packed/utils"
	"go-zone/src/model"
	"go-zone/src/service"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HtmlHomeHandler(c *gin.Context) {
	books, err := service.BookListGet()
	if err != nil {
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
	pagers := []float64{}
	len := math.Ceil(float64(total / params.Size))
	for i := 0.0; i < len; i++ {
		pagers = append(pagers, i+1)
	}
	fmt.Println("ppppppp", pagers)
	//render with master
	c.HTML(http.StatusOK, "index", gin.H{
		"title":  "Index title!",
		"books":  books,
		"posts":  list,
		"total":  total,
		"params": params,
		"pagers": pagers,
	})
}
