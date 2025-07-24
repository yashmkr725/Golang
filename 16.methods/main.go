package main

import "fmt"

// key an eye on upper case letter
func main() {
	fmt.Println("Struct in golang")
	yash := User{"Yash", "yash@gmail.com", 21, true}
	fmt.Println(yash)                             
	fmt.Printf("yash details are %+v \n", yash)   
	fmt.Printf("yash name are %v\n", yash.Username) 
	yash.GetStatus()
}

type User struct {
	Username string
	Email    string
	Age      int
	Status   bool
}

// syntax of method
func (u User) GetStatus(){ // it is working on a copy right now
	fmt.Println("Is User Active", u.Status)
} 