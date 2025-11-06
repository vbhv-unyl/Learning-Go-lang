package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://myproject.com:777/todos?course=reactjs&id=123455"

func main() {
	fmt.Println("Welcome to handling URLs in golang")

	// parsing

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of the query is %T\n", qparams)

	fmt.Println(qparams["course"])
	fmt.Println(qparams["id"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/search",
		RawPath: "user=Vaibhav",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
