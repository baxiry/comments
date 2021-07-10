package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)



func insertUser(user, pass, email string) error {
	insert, err := db.Query(
		"INSERT INTO comments.users(username, password, email) VALUES ( ?, ?, ?)",
		user, pass, email)

	// if there is an error inserting, handle it
	if err != nil {
        fmt.Println(err)
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return nil
}

func setSession(c echo.Context, username string, userid int) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60, // = 1h,
		HttpOnly: true,    // no websocket or any thing else
	}
	sess.Values["username"] = username
	sess.Values["userid"] = userid
	sess.Save(c.Request(), c.Response())
}



func login(c echo.Context) error {
	femail := c.FormValue("email")
	fpass := c.FormValue("password")
	userid, username, email, pass := getUsername(femail)

	if pass == fpass && femail == email {
		//userSession[email] = username
		setSession(c, username, userid)
		return c.Redirect(http.StatusSeeOther, "/") // 303 code
		// TODO redirect to latest page
	}
	return c.Render(200, "login.html", "Username or password is wrong")
}


func signup(c echo.Context) error {
	username := c.FormValue("username")
	pass := c.FormValue("password")
	email := c.FormValue("email")
	err := insertUser(username, pass, email)
	if err != nil {
		//fmt.Println(err)
		return c.Render(200, "sign.html", "wrrone")
	}
	return c.Redirect(http.StatusSeeOther, "/login") // 303 code
}

func signPage(c echo.Context) error {
	return c.Render(200, "sign.html", "hello")
}


func loginPage(c echo.Context) error {
	return c.Render(200, "login.html", "hello")
}

