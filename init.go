// initioal this app
package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func setdb() *sql.DB {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {

		log.Println("open database error: ", err)
		switch {
		case strings.Contains(err.Error(), "connection refused"):
			// TODO handle errors by code of error not by strings.

			//cmd := exec.Command("mysql.server", "restart")
			// for systemd linux : exec.Command("sudo", "service", "mariadb", "start")
			//cmd.Stdin = strings.NewReader(os.Getenv("JAWAD"))
			//err = cmd.Run()
			if err != nil {
				fmt.Println("error when run database cmd ", err)
			}
		default:
			log.Println("not knowen err at db.Ping() func")
			log.Println("unknown this error", err)
			os.Exit(1)
			//return nil
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
	// TODO what wrong with go:embed ?

	p := ""
	home := os.Getenv("HOME")
	if home != "/Users/fedora" {
		p = "/root/"
	}

	files := []string{
		p + "tmpl/home.html", p + "tmpl/upacount.html", p + "tmpl/acount.html", p + "tmpl/masroq.html",
		p + "tmpl/login.html", p + "tmpl/sign.html", p + "tmpl/404.html", p + "tmpl/upphotos.html",
		p + "tmpl/upcomment.html", p + "tmpl/comment.html", p + "tmpl/notfound.html", "tmpl/post.html",
		p + "tmpl/upload.html", p + "tmpl/part/header.html", p + "tmpl/part/footer.html",
	}
	return &Template{templates: template.Must(template.ParseFiles(files...))}
}

//  assets return path assets.
func assets() string {
	if os.Getenv("HOME") != "/Users/fedora" {
		return "/root/comments/assets"
	}
	return "assets"
}

/*
//  get path of photo folder
func photoFolder() string {
	if os.Getenv("HOME") == "/Users/fedora" {
		return "/Users/fedor/repo/files/"
	}
	return "../files/" // or "/root/files/"
}
*/
