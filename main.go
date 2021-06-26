package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()

    //e.Renderer = t
    e.Renderer = templ()

// start

    db := setdb()
    defer db.Close()
    
    e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))


    // files
    e.Static("/a", assets())
    e.Static("/fs", photoFold())

    e.GET("/home", homePage)//index)
    e.GET("/", index)
    e.POST("/", saveComment)
    
    // account and verefy
    e.GET("/sign", signPage)
    e.POST("/sign", signup)
    e.GET("/login", loginPage)
    e.POST("/login", login)
    e.GET("/acount/:id", acount)
    e.GET("/upacount",updateAcount)
    e.POST("/upacount",updateAcountInfo)
 
// end




	e.Logger.Fatal(e.Start(":1323"))
}

