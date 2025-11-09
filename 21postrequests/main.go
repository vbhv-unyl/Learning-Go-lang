package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to get requests in golang")

	PerformPostRequest("https://learngolang.free.beeceptor.com")
}

func PerformPostRequest(myUrl string) {

	requestBody := strings.NewReader(`
		{
			"Name": "Vaibhav",
			"Age": "22",
			"Course": "B.Tech"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseBuilder strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	byteCount, _ := responseBuilder.Write(content)
	fmt.Println("Bytecount is: ", byteCount)

	fmt.Println(responseBuilder.String())
}
