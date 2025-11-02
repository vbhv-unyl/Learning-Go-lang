package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "This is the content"

	file, err := os.Create("./sample.txt")
	if err != nil {
		// panic is used to stop the flow of program and return errors
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is : ", length)
	defer file.Close()
}
