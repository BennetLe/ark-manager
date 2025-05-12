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
	userAuth.Use(UserAuth(), CheckAdmin())

	// children := []templ.Component{
	// 	view.Index(),
	// }
	// c := layout.Base(partial.Sidebar(), children...)
	// userAuth.GET("/", CheckAdmin(), gin.WrapH(templ.Handler(c)))
	userAuth.GET("/", dashboardHandler)

	router.GET("/login", gin.WrapH(templ.Handler(layout.Login())))

	userAuth.GET("/foo", gin.WrapH(templ.Handler(partial.Foo())))

	router.POST("/login", loginHandler)
	router.GET("/logout", logoutHandler)

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

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get("user")

		// exit early if no user exists
		if user == nil {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		if !userIsAdmin(user.(string)) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}

func CheckAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get("user")

		if userIsAdmin(user.(string)) {
			ctx.Set("isAdmin", true)
			print("Admin check true")
		} else {
			ctx.Set("isAdmin", false)
			print("Admin check flase")
			print(user)
		}
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

func logoutHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusFound, "/")
}

func dashboardHandler(ctx *gin.Context) {
	isAdmin := ctx.GetBool("isAdmin")

	children := []templ.Component{}

	if isAdmin {
		children = append(children, view.Index())
	} else {
		print("No Admin")
	}
	c := layout.Base(partial.Sidebar(), children...)
	c.Render(ctx.Request.Context(), ctx.Writer)
}

func userIsAdmin(username string) bool {
	if username == "admin" {
		return true
	}
	return false
}
