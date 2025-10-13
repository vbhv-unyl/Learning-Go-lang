package main

import "fmt"

func main() {
	fmt.Println("This is ifelse")

	loginCount := 1

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
