package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)


func signPage(c echo.Context) error {
	return c.Render(200, "sign.html", "hello")
}

func loginPage(c echo.Context) error {
	return c.Render(200, "login.html", "hello")
}

// notFoundPage
func notFoundPage(c echo.Context) error {
    return c.Render(200, "notfound.html", nil)
}

func homePage(c echo.Context) error {

	sess, _ := session.Get("session", c)
	name := sess.Values["name"]
    uid := sess.Values["userid"]
	//fmt.Println("name is : ", name)

	data := make(map[string]interface{}, 3)
	data["name"] = name
    data["userid"] = uid
	data["catigories"] = catigories
    return c.Render(http.StatusOK, "home.html", data)
}

/*
// updateFotosPage router fo update Fotos Page
func updateFotosPage(c echo.Context) error {
	data := make(map[string]interface{})
    sess, _ := session.Get("session", c) // TODO i need session ?
    data["name"] = sess.Values["name"] // TODO use user id instead name
    if data["name"] == nil {
        fmt.Println("session name is nil redirect to login")
        c.Redirect(303, "/login")
    }
	
    pid := c.Param("id") 
    productId, _ := strconv.Atoi(pid)

    data["productFotos"] , err = getProductFotos(productId)
    data["userid"] = productId
    fmt.Printf("%#v", data["product"])
    if err != nil {
        fmt.Println(err)
    }
    return c.Render(http.StatusOK, "updatefotos.html", data)
}

// TODO redirect to latest page after login.
func updateProdPage(c echo.Context) error {
	// TODO whish is beter all data of product or jast photo ?
	data := make(map[string]interface{})
	sess, _ := session.Get("session", c)
	data["name"] = sess.Values["name"]
    data["userid"] = sess.Values["userid"]
	// User ID from path `users/:id`
	pid := c.Param("id") // TODO home or catigory.html ?
    productId, _ := strconv.Atoi(pid)

    fmt.Println("product id from url Param: ", productId)
	data["product"] , err = getProduct(productId)
    fmt.Printf("%#v", data["product"])
    if err != nil {
        fmt.Println(err)
    }
    return c.Render(http.StatusOK, "updateProd.html", data)
}


// upload photos
func uploadPage(c echo.Context) error {
	data := make(map[string]interface{}, 3)
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("erro upload session is : ", err)
	}
    userid := sess.Values["userid"]
	name := sess.Values["name"]
	data["name"] = name
    data["userid"] = userid
    if userid == nil {
		// TODO flash here
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	// c.Response().Status
	return c.Render(200, "upload.html", data)
}

// TODO handle error
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
    errorPage := fmt.Sprint("/404.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
    fmt.Println(err)
    //c.Redirect(303, "notfound.html")
    c.Redirect(http.StatusSeeOther, "/notfound") // 303 code
    return
}
*/

