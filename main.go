package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
    err :=  c.Render(http.StatusOK, "index.html", "Adamin")
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return nil
}

func main() {
	e := echo.New()

    e.Renderer = t
    e.GET("/", Hello)

	e.Logger.Fatal(e.Start(":1323"))
}

