package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const URL = "https://en.wikipedia.org/wiki/Go_(programming_language)"

// program to get the html of the page 

func main() {
	fmt.Println("Wikipedia Request")

	res, err := http.Get(URL)

	errCheck(err)
	fmt.Printf("the type of response is %T\n", res)

	databytes, err := io.ReadAll(res.Body)
	errCheck(err)

	content := string(databytes) 
	

	//extra
	file ,err := os.Create("./go.txt")
	errCheck(err)
	_ , err = io.WriteString(file,content)
	errCheck(err)

	fmt.Println(content)

	defer file.Close()
	defer res.Body.Close() // callers responsiblity to close the connection
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
