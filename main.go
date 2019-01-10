package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Post .
type Post struct {
	ID      int
	Content string
	Author  string
}

func main() {
	// Writing
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{ID: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{ID: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// Reading from a CSV file
	csvFileRead, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFileRead.Close()

	reader := csv.NewReader(csvFileRead)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var posts []Post
	for _, record := range records {
		id, _ := strconv.ParseInt(record[0], 0, 0)
		post := Post{
			ID:      int(id),
			Content: record[1],
			Author:  record[2],
		}
		posts = append(posts, post)
		fmt.Printf("%d\n", post.ID)
		fmt.Printf("%s\n", post.Author)
		fmt.Printf("%s\n", post.Content)
		fmt.Printf("\n")
	}

	fmt.Printf("\n\nPosts:%v\n", posts)
}
