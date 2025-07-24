package main

import "fmt"

// Generally used when everything happen then perform this task like clossing the file 

func main() {

	// defer fmt.Println("One")
	// defer fmt.Println("Two")
	// defer fmt.Println("Three")
	// fmt.Println("Defer class")
	//op -- LIFO type output
	// 	Defer class
	// Three
	// Two
	// One

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Defer class")
	mydefer()

	//OP
	// 	Defer class
	//4
	//3
	//2
	//1
	//0
	// Three
	// Two
	// One
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // it is not going to print immedeatly it value is stored instack 
	}
}
