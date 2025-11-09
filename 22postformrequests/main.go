package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to sending form data in go lang")
	PerformPostFormRequests("http://localhost:3000/postform")
}

func PerformPostFormRequests(myUrl string) {
	// Form data

	data := url.Values{}

	data.Add("Name", "Vaibhav")
	data.Add("Age", "22")
	data.Add("Email", "vaibhav.go@learn.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
