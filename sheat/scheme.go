package main

import (
	"database/sql"
	"fmt"
	"os"
)
var (
    db  *sql.DB
    err error
)

var update = `UPDATE comment SET commentor_name = 'hamed', column2 = value WHERE id = 1`
var insert = `insert into comment(link, parent_id, commentor_name, comment_text) values('localhost:1323', 1, 'omar', 'go hell comment ')`

type Comment struct {
    link, commentor_name, comment_text string
    parent_id int
}

// gets all user information for update this info
func selectComments(post_id int) Comment {
    c := Comment{}
    
    err := db.QueryRow(
        "SELECT username, email,phon, linkavatar FROM stores.users WHERE id = ?",
        post_id).Scan(&c.link, &c.parent_id, &c.commentor_name,  &c.comment_text, )
    if err != nil {
        fmt.Println("no result or", err.Error())
    }
    return c
}

func updateUserInfo(name, email, phon string, uid int) error {

    //Update db
    stmt, err := db.Prepare("update  stores.users set username=?, email=?, phon=? where id=?")
    if err != nil {return err}
    defer stmt.Close()
     
    // execute
    res, err := stmt.Exec(name, email, phon, uid)
    if err != nil {return err}
     
    a, err := res.RowsAffected()
    if err != nil {return err}
     
    fmt.Println("efected foto update: ", a)   // 1 
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
        // TODO handle this error: dial tcp 127.0.0.1:3306: connect: connection refused
        fmt.Println("mybe database is not runing or error is: ", err)
        os.Exit(1)
    }
    return db
}

