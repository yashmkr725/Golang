package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/todos/1?course=react"

func main() {
	fmt.Println("Handlings url in goLang")

	//parsing
	// Handling URL
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) //no port 
	fmt.Println(result.RawQuery) //params

	qparams := result.Query() // give key value pair 
	fmt.Println(qparams)

	// Constructing url 
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "yash.dev",
		Path: "/todos",
		RawPath: "user=yash",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
