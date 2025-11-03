package main

import "fmt"

func main() {

	/*
		Defer is a keyword which blocks the execution of the statement with which it is called
		until the caller method exits either due to ending of its scope or returns a value.

		Also the execution of multiple deferred statements follows LIFO order.
		The last deferred statement is executed first and the first last.
	*/

	defer fmt.Println("World")
	defer fmt.Println("One")
	fmt.Println("Hello")

	// Without defer output is = World One Hello
	// With defer output is = Hello One World

	myDefer()
	fmt.Println()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
