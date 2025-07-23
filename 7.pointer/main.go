package main

import "fmt"

// pointers give us the surity that whatever the resource you are passing on as a pointer instead of passing on these variables i am passing the memory address of these variables so the actual value is passed

// Sometime that you pass on these variable the copy of these varibale is being passed on whever there is a case you want to avoid such things to happen you use pointers

func main() {
	fmt.Println("Welcome to pointer class")
	//creation
	// var ptr *int ;
	// fmt.Println(ptr) // defualt value <nil>
	num := 23
	var ptr = &num // refers to 23
	fmt.Println(ptr) //0x14000112020
	fmt.Println(*ptr) //23 give the value inside

	*ptr *= 2
	fmt.Printf("the actual value is %d",num) // the actual value change
}
