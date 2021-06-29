package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)


func comment(c echo.Context) error {
    data := make(map[string]interface{},2)
    sess, _ := session.Get("session", c)
    data["userid"] = sess.Values["userid"]
    data["username"] = sess.Values["name"]

    data["comments"] = getComments("localhost:1323")

    err :=  c.Render(http.StatusOK, "comment.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;
}


func blog(c echo.Context) error {
    err :=  c.Render(http.StatusOK, "blog.html", nil)
    if err != nil {fmt.Println(err); return nil}; return nil;
}

func saveComment(c echo.Context) error {
    name := c.FormValue("name")
    comment := c.FormValue("comment")


    fmt.Println("name: ",name, "\ncomment", comment)
    return c.Render(http.StatusOK, "comment.html", "admin")
}

