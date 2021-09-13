package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Post struct {
	Post, Title, Timestamp string
	PostId, Ownerid        int
}

// got  post by postid from database
func getPost(postid int) (Post, error) {
	p := Post{}
	err := db.QueryRow(
		"SELECT title, post,  creatAt, ownerId FROM comments.posts WHERE postId = ?",
		postid).Scan(&p.Title, &p.Post, &p.Timestamp, &p.Ownerid)

	if err != nil {
		return p, err
	}
	fmt.Println("post in db function si: ", p)
	return p, nil
}

func showPost(postid int, c echo.Context) error {
	data := make(map[string]interface{}, 1)

	data["post"], err = getPost(1)
	if err != nil {
		return err
	}

	fmt.Println("data in showPost is :", data["post"])
	// return c.Render(http.StatusOK, "comment.html", data)
	err := c.Render(http.StatusOK, "post.html", data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

// getPosts get all posts titles from database
func getPosts() []Post {
	// get just titles of all posts
	qur := "select ownerid, postText, timestamp from comments.posts"
	rows, err := db.Query(qur)
	if err != nil {
		fmt.Println("at query func owner id db select ", err)
	}
	defer rows.Close() // ??

	posts := make([]Post, 1)
	p := Post{}

	// iterate over rows
	for rows.Next() {
		err = rows.Scan(&p.PostId, &p.Post, &p.Ownerid, &p.Timestamp)
		if err != nil {
			fmt.Println("error At getPosts function", err)
		}
		posts = append(posts, p)

	}
	fmt.Println("lenght of data is : ", len(posts))
	return posts
}

// save comment in database
func savePost(c echo.Context) error {

	sess, _ := session.Get("session", c)
	data := make(map[string]interface{}, 2)
	data["userid"] = sess.Values["userid"]
	data["username"] = sess.Values["username"]

	post := c.FormValue("post")

	// TODO save comment and get data

	fmt.Println("user id : ", data["userid"], "  post: ", post)
	// return c.Render(http.StatusOK, "comment.html", data)
	err := c.Render(http.StatusOK, "post.html", data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

// TODO update Post
func updatPost() {}

// TODO delete Post
func delletePost() {}

// render Post Page
func postPage(c echo.Context) error {

	data := make(map[string]interface{}, 2)
	sess, _ := session.Get("session", c)
	data["userid"] = sess.Values["userid"]
	data["username"] = sess.Values["username"]
	fmt.Println("usser nam is : ", data["username"])

	data["post"], err = getPost(1) //"this is a best post you can read in your life. you agree with me ? no ? then go to hell"
	if err != nil {
		fmt.Println("erro is : ", err)
	}
	fmt.Println("post is : ", data["post"])

	err := c.Render(http.StatusOK, "post.html", data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
	//return c.Render(http.StatusOK, "post.html", data)
}
