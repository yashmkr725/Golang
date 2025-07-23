package main

import "fmt"

const Logintoken string = "iamamonkey" // If the variable first letter is capital it can be acessed by any other file and you can use this anywhere
//Public

func main() {
	var username string = "Yash"
	fmt.Println(username)
	fmt.Printf("Variable is of type is %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type is %T \n", isLogged)

	var small int64 = 255
	fmt.Println(small)
	fmt.Printf("Variable is of type is %T \n", small)

	var pointing float32 = 255.5343
	fmt.Println(pointing)
	fmt.Printf("Variable is of type is %T \n", pointing)

	// default values and aliasis
	var anoterInt int
	fmt.Println(anoterInt)
	fmt.Printf("Variable is of type is %T \n", anoterInt) // default zero 

	var anotherstring string
	fmt.Println(anotherstring) // default empty
	fmt.Printf("Variable is of type is %T \n", anotherstring)

	// implicit type
	var website = "hello"
	fmt.Println(website)

	// no var style - inside of method you are allowed to use valrious operator but not outside / global
	noOfUsers := 30000
	fmt.Println(noOfUsers)
}
