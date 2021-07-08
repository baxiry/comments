package main

import (
	"database/sql"
	"fmt"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)


type Comment struct {
    CommentId, ParentId,  UserId              int
    CommentText,  Link, AvatarLink, TimeStamp string
}


func getComments(link string) []*Comment {
    qur := "select comment_id, user_id, parent_id, comment_text, created_at from comments.comments where link = ?"
    rows, err := db.Query(qur, link)
	if err != nil {
		fmt.Println("at query func owner id db select ", err)
	}
	defer rows.Close() // ??

    comments := make([]*Comment,0)
    c := Comment{}

	// iterate over rows
	for rows.Next() {
		err = rows.Scan(&c.CommentId, &c.ParentId, &c.UserId, &c.CommentText, &c.TimeStamp)
		if err != nil {
			fmt.Println("At get all my product", err)
		}
        comments = append(comments, &c)

	}
    fmt.Println("lenght of data is : ", len(comments))
	return comments
}


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




func updateUserInfo(name, email string, uid int) error {

	//Update db
	stmt, err := db.Prepare("update  comments.users set username=?, email=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	res, err := stmt.Exec(name, email, uid)
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

// gets all user information for update this info
func getUserInfo(userid int) ( string, string, string) {
	var name, email, avatar string
	err := db.QueryRow(
        "SELECT username, email, linkavatar FROM comments.users WHERE user_id = ?",
		userid).Scan(&name, &email, &avatar)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
    fmt.Println("name is : ", name, "email is : ", email, "avatar is ", avatar)
	return name, email, avatar
}

// get all username
func getUsername(femail string) (int, string, string, string) {
	var name, email, password string
	var userid int
	err := db.QueryRow(
        "SELECT user_id, username, email, password FROM comments.users WHERE email = ?",
		femail).Scan(&userid, &name, &email, &password)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return userid, name, email, password
}

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

func setdb() *sql.DB {
	db, err = sql.Open(
		"mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil { // why no error when db is not runinig ??
		fmt.Println("run mysql server", err)
		// TODO report this error.

		// wehen db is stoped no error is return.
		// we expecte errore no database is runing

		// my be this error is fixed with panic ping pong bellow
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
