package router

import (
	"go-zone/packed/utils"
	"go-zone/src/controller"
	"html/template"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
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
			"copy":      utils.CopyRight,
			"innerHtml": utils.InnerHtml,
		},
		DisableCache: true,
	})

	r.GET("/", controller.HtmlHomeHandler)

	r.GET("/p/:slue", controller.HtmlPostDetailHandler)
}
