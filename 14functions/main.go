package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(1, 2)
	fmt.Println("Result is: ", result)

	proResult := proAdder(2, 4, 5, 6, 7)
	fmt.Println("Pro result is: ", proResult)

	number, message := messager()
	fmt.Println("Number is: ", number)
	fmt.Println("Message is: ", message)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func messager() (int, string) {
	return 1, "Hello"
}
