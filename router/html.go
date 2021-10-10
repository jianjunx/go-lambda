package router

import (
	"go-zone/src/controller"
	"html/template"
	"net/http"
	"time"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func registerHtml() {
	r.StaticFS("/public", http.Dir("public"))
	//new template engine
	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad", "partials/header"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	r.GET("/", controller.HtmlHomeHandler)

	r.GET("/p/:slue", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.tpl", gin.H{"title": "Page file title!!"})
	})
}
