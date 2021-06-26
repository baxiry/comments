package main

import (
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// path file is depends to enveronment.
func templ() *Template {
	var p string
	if os.Getenv("USERNAME") != "fedor" {
		p = "/root/store/"
	}
	files := []string{
        p + "tmpl/home.html",p + "tmpl/upacount.html",  p + "tmpl/acount.html", p + "tmpl/login.html", p + "tmpl/sign.html",
        p + "tmpl/404.html", p + "tmpl/updateProd.html",
        p + "tmpl/stores.html",p + "tmpl/mystore.html", p + "tmpl/notfound.html",
		p + "tmpl/upload.html", p + "tmpl/product.html", p + "tmpl/products.html",
        p + "tmpl/part/header.html", p + "tmpl/part/footer.html", p + "tmpl/updatefotos.html",

	}
	return &Template{templates: template.Must(template.ParseFiles(files...))}
}


// folder when photos is stored.
func photoFold() string {
	if os.Getenv("USERNAME") == "fedor" {
		return "/home/fedor/repo/files/"
	}
	return "/root/files/"
}


// where assets  path ?
func assets() string {
	if os.Getenv("USERNAME") != "fedor" {
		return "/root/store/assets"
	}
	return "assets"
}


var catigories = map[string][]string{
	"cars":      {"mersides", "volswagn", "shefrole", "ford", "jarary", "jawad"},
	"animals":   {"dogs", "sheeps", "elephens", "checkens", "lions"},
	"motors":    {"harly", "senteroi", "basher", "hddaf", "mobilite"},
	"mobiles":   {"sumsung", "apple", "oppo", "netro", "nokia"},
	"computers": {"dell", "toshipa", "samsung", "hwawi", "hamed"},
	"services":  {"penter", "developer", "cleaner", "shooter", "gamer"}, //services
	"others":    {"somthing", "another-somth", "else", "anythings"},
}

