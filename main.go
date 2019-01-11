package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

// Post .
type Post struct {
	ID      int
	Content string
	Author  string
}

func store(data interface{}, fileName string) {
	bytesBuffer := new(bytes.Buffer)
	bytesEncoder := gob.NewEncoder(bytesBuffer)
	err := bytesEncoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, bytesBuffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {

	allPosts := []Post{
		Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{ID: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{ID: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	store(allPosts, "gob_file_1")

}
