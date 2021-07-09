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
    data["username"] = sess.Values["username"]
    fmt.Println( "usser nam is : ", data["username"])

    data["comments"] = getComments("localhost:1323") // get comments by link of article

    err :=  c.Render(http.StatusOK, "comment.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;
}


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

func saveComment(c echo.Context) error {
    
    sess, _ := session.Get("session", c)
    data := make(map[string]interface{}, 2)
    data["userid"] = sess.Values["userid"]
    data["username"] = sess.Values["username"]

    name := c.FormValue("name")
    comment := c.FormValue("comment")

    // TODO save comment and get data

    fmt.Println("name: ",name, "\ncomment", comment)
    // return c.Render(http.StatusOK, "comment.html", data)
    err :=  c.Render(http.StatusOK, "comment.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;

}

