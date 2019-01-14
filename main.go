package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Post .
type Post struct {
	ID      int
	Content string
	Author  string
}

// Db .
var Db *sql.DB

func init() {
	var err error
	connStr := "user=gwp password=gwp dbname=gwp sslmode=disable host=localhost port=5432"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

// Create .
func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// GetPost .
func GetPost(id int) (post Post, err error) {
	post = Post{}
	Db.QueryRow("SELECT id, content, author FROM posts WHERE id=$1", id).Scan(&post.ID, &post.Content, &post.Author)
	return post, nil
}

// Update .
func (post *Post) Update() (err error) {
	result, err := Db.Exec("update posts set content=$2, author=$3 where id=$1", post.ID, post.Content, post.Author)
	fmt.Printf("\n\nResult returned after the update: %v\n\n", result)
	return
}

// Delete .
func (post *Post) Delete() (err error) {
	result, err := Db.Exec("delete from posts where id = $1", post.ID)
	fmt.Printf("\n\nResult returned after the delete: %v\n\n", result)
	return
}

func main() {
	var err error
	// Create a post
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)
	err = post.Create()
	if err != nil {
		fmt.Printf("Error inserting post: %v\n", err)
	}
	fmt.Println(post)

	// Retrieve a post
	readPost, err := GetPost(1)
	fmt.Printf("Post read: %v\n\n", readPost)

	// Update the post
	readPost.Content = "Once up on a time in China"
	readPost.Author = "Huang Ho"
	readPost.Update()

	// Delete the post
	readPost.Delete()
}
