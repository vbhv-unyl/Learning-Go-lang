package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to get requests in golang")

	PerformGetRequest("https://jsonplaceholder.typicode.com/todos/1")
}

func PerformGetRequest(myUrl string) {
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseBuilder strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

	byteCount, _ := responseBuilder.Write(content)
	fmt.Println("Bytecount is: ", byteCount)

	fmt.Println(responseBuilder.String())
}
