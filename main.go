package main

import (
	"ark-manager/view"
	"ark-manager/view/layout"
	"ark-manager/view/partial"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static/", "./static")

	c := layout.Base(view.Index())
	router.GET("/", gin.WrapH(templ.Handler(c)))

	router.GET("/foo", gin.WrapH(templ.Handler(partial.Foo())))

	router.Run(":8080")
}
