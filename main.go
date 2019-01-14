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
	} else {
		fmt.Printf("\n\n%v\n\n", Db)
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

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)
	err := post.Create()
	if err != nil {
		fmt.Printf("Error inserting post: %v\n", err)
	}
	fmt.Println(post)
}
