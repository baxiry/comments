package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", "Adamin")
    // if err != nil {fmt.Println(err); return nil}; return nil;
}



func saveComment(c echo.Context) error {
    name := c.FormValue("name")
    comment := c.FormValue("comment")
    fmt.Println("name: ",name, "\ncomment", comment)
    return c.Render(http.StatusOK, "index.html", "admin")
}


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

    e.GET("/", homePage)//index)
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

