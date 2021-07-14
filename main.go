package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
    e.Renderer = templ()
    db := setdb()
    defer db.Close()
   

    // middleware
    e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

    // routers
    e.GET("/api", commentsPage)
    e.POST("/api", saveComment)

    e.GET("/", homePage)// index
    e.GET("/blog", blog)
    e.GET("/post", postPage)
    
    // account and verefy
    e.GET("/sign", signPage)
    e.POST("/sign", signup)
    e.GET("/login", loginPage)
    e.POST("/login", login)
    e.GET("/acount/:id", acount)
    e.GET("/upacount",updateAcount)
    e.POST("/upacount",updateAcountInfo)
 
    // files
    e.Static("/a", assets())
    e.Static("/fs", photoFold())

	e.Logger.Fatal(e.Start(":1323"))
}

