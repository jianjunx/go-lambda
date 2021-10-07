package controller

import (
	"context"
	"go-zone/src/model"
	"go-zone/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addPostBody struct {
	Data model.PostAttributes `json:"data"`
}

func PostWebhookYuqueHandler(c *gin.Context) {
	p := addPostBody{}
	context.TODO()
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = service.PostItemDataAdd(&p.Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
