package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

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

}
