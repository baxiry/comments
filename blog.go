package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)
func blog(c echo.Context) error {
    data := make(map[string]interface{}, 2)
    
    sess, _ := session.Get("session", c)
    data["userid"] = sess.Values["userid"]
    data["username"] = sess.Values["username"]
    fmt.Println("user name is : ", data["username"])
   
    // return c.Render(http.StatusOK, "comment.html", data)
    err :=  c.Render(http.StatusOK, "blog.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;
}


