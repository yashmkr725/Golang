package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // we use this valrous operator becuase we dont knw which type of value is goind to come from package
	fmt.Println("Entre the reading for pizza")

	// comma ok || error ok - it says either you get an input or recieve an error 
	// just like try catch
	input , _ :=  reader.ReadString('\n') // in go we use _ for value we dont care about
	fmt.Println("Thanks for rating",input)
	fmt.Printf("Type of the rating is %T",input)
}
