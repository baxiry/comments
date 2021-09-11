package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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
		fmt.Println("error when open mysql server", err)
		// TODO report this error.
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("error when ping to database", err)
		switch {
		case strings.Contains(err.Error(), "connection refused"):
			// TODO handle errors by code of error not by strings.

			cmd := exec.Command("mysql.server", "restart") // for systemd linux : exec.Command("sudo", "service", "mariadb", "start")
			//cmd.Stdin = strings.NewReader(os.Getenv("JAWAD"))
			errc := cmd.Run()
			if errc != nil {
				fmt.Println("error when run database cmd ", errc)
			}
		default:
			fmt.Println("error at  setdb() func, db.Ping() func")
			fmt.Println("unknown this error", err)
			os.Exit(1)
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
	// TODO use go:embed

	files := []string{
		"tmpl/home.html", "tmpl/upacount.html", "tmpl/acount.html",
		"tmpl/login.html", "tmpl/sign.html", "tmpl/404.html", "tmpl/upphotos.html",
		"tmpl/upcomment.html", "tmpl/comment.html", "tmpl/notfound.html", "tmpl/post.html",
		"tmpl/upload.html", "tmpl/part/header.html", "tmpl/part/footer.html",
	}
	return &Template{templates: template.Must(template.ParseFiles(files...))}
}

/*
//  get path of photo folder
func photoFold() string {
	//if os.Getenv("USERNAME") == "fedor" {
	//	return "/home/fedor/repo/files/"
	//}
	return "../files/" // or "/root/files/"
}

// TODO use go:embed for assets

//  assets return path assets.
func assets() string {
	//if os.Getenv("USERNAME") != "fedor" {
	//	return "/root/store/assets"
	//}
	fmt.Println("we are on mac")
	return "assets"
}
*/
