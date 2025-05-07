package main

import (
	"ark-manager/view"
	"ark-manager/view/layout"
	"ark-manager/view/partial"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static/", "./static")

	c := layout.Base(view.Index())
	router.GET("/", gin.WrapH(templ.Handler(c)))

	router.GET("/login", gin.WrapH(templ.Handler(layout.Login())))
	router.GET("/foo", gin.WrapH(templ.Handler(partial.Foo())))

	router.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"uname": username,
			"pswd":  password,
		})
	})

	router.Run(":8080")
}
