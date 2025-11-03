package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("Length is : ", length)
	defer file.Close()

	readFile("./sample.txt")
}

func readFile(filename string) {
	// ioutil is deprecated. Look for what is used inplace of it now.
	databyte, err := ioutil.ReadFile(filename)
	// if(err != nil) {
	// 	panic(err)
	// }

	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
