package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Comment struct {
    CommentId, ParentId,  UserId              int
    CommentText,  Link, AvatarLink, TimeStamp string
}


// render comments
func commentsPage(c echo.Context) error {

    data := make(map[string]interface{},2)
    sess, _ := session.Get("session", c)
    data["userid"] = sess.Values["userid"]
    data["username"] = sess.Values["username"]
    fmt.Println( "usser nam is : ", data["username"])

    data["comments"] = getComments("localhost:1323") // get comments by link of article

    err :=  c.Render(http.StatusOK, "comment.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;
}

// get comments of article from database
func getComments(link string) []Comment {
    qur := "select comment_id, user_id, parent_id, comment_text, created_at from comments.comments where link = ?"
    rows, err := db.Query(qur, link)
	if err != nil {
		fmt.Println("at query func owner id db select ", err)
	}
	defer rows.Close() // ??

    comments := make([]Comment,0)
    c := Comment{}

	// iterate over rows
	for rows.Next() {
        err = rows.Scan(&c.CommentId, &c.UserId, &c.ParentId, &c.CommentText, &c.TimeStamp)
		if err != nil {
			fmt.Println("At get all my product", err)
		}
        comments = append(comments, c)

	}
    fmt.Println("lenght of data is : ", len(comments))
	return comments
}


// save comment in database
func saveComment(c echo.Context) error {
    
    sess, _ := session.Get("session", c)
    data := make(map[string]interface{}, 2)
    userid := sess.Values["userid"]
    data["username"] = sess.Values["username"]
    comment := c.FormValue("comment")
    parentid := c.QueryParam("parentid")


    // TODO save comment and get data

    fmt.Println( "user id", userid,"  comment", comment, parentid)
    // return c.Render(http.StatusOK, "comment.html", data)
    err :=  c.Render(http.StatusOK, "comment.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;
}

/*

func updateProductFotos(photos string, id int) error {

	//Update db
	stmt, err := db.Prepare("update  comments.products set photos=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	res, err := stmt.Exec(photos, id)
	if err != nil {
		return err
	}

	a, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("efected foto update: ", a) // 1
	return nil
}

func updateProduct(title, catig, descr, price, photos string, id int) error {

	//Update db
	stmt, err := db.Prepare("update  comments.products set  title=?,  catigory=?, description=?,  price=?,  photos=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	res, err := stmt.Exec(title, catig, descr, price, photos, id)
	if err != nil {
		return err
	}

	a, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(a) // 1
	return nil
}

// delete Producte.
func deleteProducte(id int) error {
	res, err := db.Exec("DELETE FROM comments.products WHERE id=?", id)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		return err
	}
	fmt.Println("affectedRows: ", affectedRows)
	// defer res // TODO I need understand this close in mariadb
	return nil
}


func getProduct(id int) (Product, error) {
	var p Product
	var picts string
	err := db.QueryRow(
		"SELECT title, catigory, description, photos, price FROM comments.products WHERE id = ?",
		id).Scan(&p.Title, &p.Catigory, &p.Description, &picts, &p.Price)
	if err != nil {
		return p, err
	}

	list := strings.Split(picts, "];[")
	fmt.Println("list fotos is :", list)
	// TODO split return 2 item in some casess, is this a bug ?
	p.Photos = filter(list)
	p.Id = id
	return p, nil
}

// getCatigories get all photo name of catigories.
func getProductes(catigory string) ([]Product, error) {
	var p Product
	var picts string
	res, err := db.Query(
		"SELECT id, title, photos, price FROM comments.products WHERE catigory = ?", catigory)
	if err != nil {
		return nil, err
	}
	defer res.Close() // TODO I need understand this close in mariadb

	items := make([]Product, 0)
	for res.Next() {
		res.Scan(&p.Id, &p.Title, &picts, &p.Price)
		list := strings.Split(picts, "];[")
		// TODO split return 2 item in some casess, is this a bug ?
		p.Photo = list[0]
		items = append(items, p)
		// TODO we need just avatar photo
	}
	return items, nil
}

func insertProduct(title, catigory, details, picts string, ownerid, price int) error {
	insert, err := db.Query(
		"INSERT INTO comments.products(ownerid, title, catigory, description, price, photos) VALUES ( ?, ?, ?, ?, ?, ?)",
		ownerid, title, catigory, details, price, picts)
	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close() // TODO why we need closeing this connection ?

	return nil
}
// some tools
func filter(slc []string) []string {
	res := make([]string, 0)
	for _, v := range slc {
		if v != "" {
			res = append(res, v) // TODO this need improve fo performence
		}
	}
	return res
}

type Product struct {
	Id          int
	Title       string
	Catigory    string
	Description string
	Photo       string
	Photos      []string
	Price       string
}

*/
