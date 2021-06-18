package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)



var t = &Template{
    templates: template.Must(template.ParseFiles("./index.html")),
}

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}


