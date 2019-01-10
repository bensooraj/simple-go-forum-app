package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello, world!\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("data1")
	fmt.Printf("%s", read1)

	file1, _ := os.Create("data2")
	defer file1.Close()

	b, _ := file1.Write(data)
	fmt.Printf("Wrote %v bytes.\n", b)

	file2, _ := os.Open("data2")
	defer file2.Close()

	chunk := make([]byte, len(data))
	b, _ = file2.Read(chunk)

	fmt.Printf("Read %d bytes from file\n", b)
	fmt.Println(string(chunk))
}
