package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// this is used to make get request on another server
func main() {
	fmt.Println("Get method in Go")
	PerfromGetRequest()
}

func PerfromGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	res, err := http.Get(myurl)
	checkError(err)

	defer res.Body.Close() // all things done it is going to close this

	fmt.Println("Status code", res.Status)
	fmt.Println("Content Length", res.ContentLength)

	var responseString strings.Builder
	content, err := io.ReadAll(res.Body) // content is in byte format
	checkError(err)
	byteCount, _ := responseString.Write((content))
	fmt.Println(byteCount) // 83

	// fmt.Println(string(content)) another way using responseString
	fmt.Println(responseString.String()) // better way it uses a library more powerful more features

	// OP
	// 	Get method in Go
	// Status code 200 OK
	// Content Length -1
	// 83
	// {
	//   "userId": 1,
	//   "id": 1,
	//   "title": "delectus aut autem",
	//   "completed": false
	// }

}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
