package main

import (
	"fmt"
)
// we dont use array much in golang we have another dtype slices
func main() {
	fmt.Println("Welcome to array in golang")
	var fruits [4]string // while declaring an  array it is imp to mension size 
	fruits[0] = "apple"
	fruits[1] = "mango"
	fruits[3] = "banana"
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var veges =[3]string{"potato","carrot","tomato"}
	fmt.Println(veges[0])
}
