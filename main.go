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

func initDB() {
	connStr := "user=gwp dbname=gwp sslmode=disable host=localhost port=5432"
	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\n\n%v\n\n", Db)
	}

}

func main() {
	initDB()
}
