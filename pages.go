package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func homePage(c echo.Context) error {

	sess, _ := session.Get("session", c)

	data := make(map[string]interface{}, 3)
	data["username"] = sess.Values["username"]
	data["userid"] = sess.Values["userid"]

	return c.Render(http.StatusOK, "home.html", data)
}

func masroq(c echo.Context) error {
	return c.Render(http.StatusOK, "masroq.html", nil)
}

/*


// notFoundPage
func notFoundPage(c echo.Context) error {
    return c.Render(200, "notfound.html", nil)
}



// updateFotosPage router fo update Fotos Page
func updateFotosPage(c echo.Context) error {
	data := make(map[string]interface{})
    sess, _ := session.Get("session", c)
    data["name"] = sess.Values["name"]
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

func updateProdPage(c echo.Context) error {
	data := make(map[string]interface{})
	sess, _ := session.Get("session", c)
	data["name"] = sess.Values["name"]
    data["userid"] = sess.Values["userid"]
	// User ID from path `users/:id`
	pid := c.Param("id")
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
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	// c.Response().Status
	return c.Render(200, "upload.html", data)
}

*/
