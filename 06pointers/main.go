package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to pointers study of golang!")

	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr)

	myNumber := 23
	var ptr *int = &myNumber

	fmt.Println("Value of pointer is:", ptr)
	fmt.Println("Value of number is:", *ptr)
}
