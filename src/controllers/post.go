package controllers

import (
	"context"
	"gin-template/src/model"
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
		return
	}
	c.JSON(http.StatusOK, nil)
}
