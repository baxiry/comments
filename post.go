package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)


type Post struct {
    postId, UserId              int
    post,  avatarLink, simeStamp string
}


// render Post Page 
func postPage(c echo.Context) error {

    data := make(map[string]interface{},2)
    sess, _ := session.Get("session", c)
    data["userid"] = sess.Values["userid"]
    data["username"] = sess.Values["username"]
    fmt.Println( "usser nam is : ", data["username"])

    data["post"] = "this is a best post you can read in your life. you agree with me ? no ? then go to hell"

    err :=  c.Render(http.StatusOK, "post.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;
}

// get Posts from database
func getPosts(link string) []Comment {
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
func savePost(c echo.Context) error {
    
    sess, _ := session.Get("session", c)
    data := make(map[string]interface{}, 2)
    data["userid"] = sess.Values["userid"]
    data["username"] = sess.Values["username"]

    comment := c.FormValue("comment")

    // TODO save comment and get data

    fmt.Println( "user id", data["userid"],"  comment", comment)
    // return c.Render(http.StatusOK, "comment.html", data)
    err :=  c.Render(http.StatusOK, "comment.html", data)
    if err != nil {fmt.Println(err); return nil}; return nil;
}

func updatPost(){}
func delletePost(){}
