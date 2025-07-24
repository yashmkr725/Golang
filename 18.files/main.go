package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "Hello this is the text goes in file"

	file, err := os.Create("./a.txt")

	checkNilError((err))

	fmt.Println(file)

	length, err := io.WriteString(file, content)

	checkNilError((err))

	fmt.Println(length)

	defer file.Close()
	ReadFile("./a.txt")
}

func ReadFile(filename string) {
	databyte, err := os.ReadFile(filename) // io.ReadFile is depreciated now we use os.Readfile
	checkNilError((err))
	fmt.Println(databyte)
	fmt.Println(string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
