package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)


func index(c echo.Context) error {
    err :=  c.Render(http.StatusOK, "index.html", "Adamin")
    if err != nil {fmt.Println(err); return nil}; return nil;
}



func saveComment(c echo.Context) error {
    name := c.FormValue("name")
    comment := c.FormValue("comment")
    fmt.Println("name: ",name, "\ncomment", comment)
    return c.Render(http.StatusOK, "index.html", "admin")
}

