package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Post .
type Post struct {
	ID         int
	Content    string
	AuthorName string `db:"author"`
	Comments   []Comment
}

// Comment .
type Comment struct {
	ID      int
	Content string
	Author  string
	Post    *Post
}

// Db .
var Db *sqlx.DB

func init() {
	var err error
	connStr := "user=gwp password=gwp dbname=gwp sslmode=disable host=localhost port=5432"
	Db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		// panic(err)
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

	err = stmt.QueryRow(post.Content, post.AuthorName).Scan(&post.ID)
	return
}

// Create .
func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post Not Found")
		return
	}

	err = Db.QueryRow("INSERT INTO comments (content, author, post_id) VALUES ($1, $2, $3) RETURNING id", comment.Content, comment.Author, comment.Post.ID).Scan(&comment.ID)

	return
}

// GetPost .
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	// Db.QueryRow("SELECT id, content, author FROM posts WHERE id=$1", id).Scan(&post.ID, &post.Content, &post.AuthorName)
	Db.QueryRowx("SELECT id, content, author FROM posts WHERE id=$1", id).StructScan(&post)

	// Retrieve comments as well
	fmt.Printf("Post retrieved: %v\n\n", post)
	rows, err := Db.Query("SELECT id, content, author FROM comments WHERE post_id=$1", post.ID)
	if err != nil {
		return
	}
	// defer rows.Close()

	for rows.Next() {
		comment := Comment{Post: &post}
		err := rows.Scan(&comment.ID, &comment.Content, &comment.Author)
		if err != nil {
			fmt.Printf("Comments row scan error: %v\n\n", err)
			return post, err
		}
		// fmt.Printf("Comment: %v\n", comment)
		post.Comments = append(post.Comments, comment)
	}

	return post, nil
}

// Update .
func (post *Post) Update() (err error) {
	result, err := Db.Exec("update posts set content=$2, author=$3 where id=$1", post.ID, post.Content, post.AuthorName)
	fmt.Printf("\n\nResult returned after the update: %v\n\n", result)
	return
}

// Delete .
func (post *Post) Delete() (err error) {
	result, err := Db.Exec("delete from posts where id = $1", post.ID)
	fmt.Printf("\n\nResult returned after the delete: %v\n\n", result)
	return
}

// Posts .
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, content, author FROM posts LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Content, &post.AuthorName)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Comments .
func Comments(limit int) (comments []Comment, err error) {
	rows, err := Db.Query("SELECT id, content, author FROM comments LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.ID, &comment.Content, &comment.Author)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go printNumbers2(&wg)
	// go printLetters2(&wg)
	// wg.Wait()
	// fmt.Println()
	// fmt.Println()

	// w1, w2 := make(chan bool), make(chan bool)
	// go printNumbers2(w1)
	// go printLetters2(w2)

	// result1 := <-w1
	// result2 := <-w2

	// if result1 && result2 {
	// 	fmt.Println("Both the results are true!")
	// }
	c := make(chan int, 2)
	go thrower(c)
	go catcher(c)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello, Ben!")

}
