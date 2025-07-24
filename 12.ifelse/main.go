package main

import "fmt"

func main() {
	fmt.Println("If Else Lecture")

	loginCount := 0
	var result string

	if loginCount < 0 {
		result = "New User"
	} else if loginCount == 0 {
		result = "Go Go"
	} else {
		result = "Old User"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("No is even")
	} else {
		fmt.Println("No is odd")
	}

	// there is a syntax where you initialise a variable and directly put up a value into it used mostly in web

	if num := 3; num < 3 { // initializing ,assigning value and check
		fmt.Println("no is less than 5")
	} else {
		fmt.Println("no is not less than 10")
	}

}
