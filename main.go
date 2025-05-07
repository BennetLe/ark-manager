package main

import (
	"ark-manager/view"
	"ark-manager/view/layout"
	"ark-manager/view/partial"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static/", "./static")

	store := cookie.NewStore([]byte("secret")) // use a secure secret in production
	router.Use(sessions.Sessions("session", store))

	userAuth := router.Group("/")
	userAuth.Use(UserAuth())

	children := []templ.Component{
		view.Index(),
	}
	c := layout.Base(partial.Sidebar(), children...)
	userAuth.GET("/", gin.WrapH(templ.Handler(c)))

	router.GET("/login", gin.WrapH(templ.Handler(layout.Login())))
	router.GET("/foo", gin.WrapH(templ.Handler(partial.Foo())))

	router.POST("/login", loginHandler)

	router.Run(":8080")
}

func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get("user")
		if user == nil {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

func loginHandler(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if password == "test" {
		session := sessions.Default(ctx)
		session.Set("user", username)
		session.Save()
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
}
