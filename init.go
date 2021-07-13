package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/exec"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var (
	db  *sql.DB
	err error
)


func setdb() *sql.DB {
	db, err = sql.Open(
		"mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil { // why no error when db is not runinig ??
		fmt.Println("run mysql server", err)
		// TODO report this error.
	}

	if err = db.Ping(); err != nil {

        cmd := exec.Command("sudo", "service", "mariadb", "start") 
        //cmd.Stdin = strings.NewReader(os.Getenv("JAWAD"))                        
        errc := cmd.Run()                                                     
        if errc != nil {                                                      
            fmt.Println(errc)                                                   
        }
	}
	return db
}


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
        p + "tmpl/home.html",p + "tmpl/upacount.html", p + "tmpl/acount.html",p + "tmpl/blog.html",
        p + "tmpl/login.html", p + "tmpl/sign.html", p + "tmpl/404.html", p + "tmpl/upphotos.html",
        p + "tmpl/upcomment.html", p + "tmpl/comment.html", p + "tmpl/notfound.html",p + "tmpl/post.html", 
		p + "tmpl/upload.html", p + "tmpl/part/header.html", p + "tmpl/part/footer.html",

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

