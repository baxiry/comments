package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// updateAcount updates Acount information
func updateAcountInfo(c echo.Context) error {
	//data := make(map[string]interface{},1)
	sess, _ := session.Get("session", c)
	uid := sess.Values["userid"]
	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	fmt.Println("username+email is :", username, email)

	err := updateUserInfo(username, email, uid.(int))
	if err != nil {
		fmt.Println("error at update db function", err)
	}

	// update session information
	setSession(c, username, uid.(int))

	// redirect to acoun page
	userid := strconv.Itoa(uid.(int))

	return c.Redirect(303, "/acount/"+userid)
}

// updateAcount updates Acount information
func updateAcount(c echo.Context) error {
	data := make(map[string]interface{}, 1)
	sess, _ := session.Get("session", c)

	userid := sess.Values["userid"]
	username := sess.Values["username"]

	data["username"] = username

    if userid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

    data["username"], data["email"], data["linkavatar"] = getUserInfo(userid.(int))

    data["userid"] = userid

	fmt.Println(data)

	return c.Render(200, "upacount.html", data)
}

// acount render profile of user.
func acount(c echo.Context) error {
	sess, _ := session.Get("session", c)
	data := make(map[string]interface{}, 2)
	data["username"] = sess.Values["username"]
    data["userid"] = sess.Values["userid"]
	// TODO get all info like foto from db

    if data["userid"] == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	return c.Render(200, "acount.html", data)
}

//
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.Render(http.StatusOK, "user.html", id)
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
