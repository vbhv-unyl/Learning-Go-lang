package main

import "fmt"

// In case of global variables, direct assignment with valrus operator ':=' is not poaaible
// You have to use the proper assignment syntax, defining the variable type is optional

var global = 100

const LoginToken string = "Text" // Public

func main() {
	var username string = "Vaibhav"
	fmt.Println(username)
	fmt.Printf("Type of variable : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of variable : %T \n", smallVal)

	var smallFloat float32 = 255.23242
	fmt.Println(smallFloat)
	fmt.Printf("Type of variable : %T \n", smallFloat)

	var temp = "String"
	fmt.Println(temp)

	t := 34
	fmt.Println(t)

	fmt.Println(LoginToken)
	fmt.Printf("Type of variable : %T \n", LoginToken)

}
