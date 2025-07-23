package main

import "fmt"

// key an eye on upper case letter
func main() {
	fmt.Println("Struct in golang")
	// There is no inhertance in golang ; no super or parent these concept is not in golang
	yash := User{"Yash", "yash@gmail.com", 21, true}
	fmt.Println(yash)                             //{Yash yash@gmail.com 21 true}
	fmt.Printf("yash details are %+v \n", yash)   //yash details are {Username:Yash Email:yash@gmail.com Age:21 Status:true}
	fmt.Printf("yash name are %v", yash.Username) //yash name are Yash
}

type User struct {
	Username string
	Email    string
	Age      int
	Status   bool
}
