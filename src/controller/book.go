package controller

import (
	"go-zone/src/service"

	"github.com/gin-gonic/gin"
)

func BookListGetHandler(c *gin.Context) {
	books, err := service.BookListGet()
	if err != nil {
		respError(c, err)
		return
	}
	respSuccess(c, books)
}
