package utils

import (
	"html/template"
	"math"
	"strconv"
	"time"
)

func StrToInt(str string, defValue int) int {
	if str == "" {
		return defValue
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return defValue
	}
	return val
}

func InnerHtml(htmlStr string) template.HTML {
	return template.HTML(htmlStr)
}

func CopyRight() string {
	return time.Now().Format("2006")
}

func GetPages(total, size float64) (pages []float64) {
	pages = []float64{}
	for i, len := 0.0, math.Ceil(total/size); i < len; i++ {
		pages = append(pages, i+1)
	}
	return
}
