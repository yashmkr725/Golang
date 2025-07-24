package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)
	// ++i is not in go

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// In most of the cases you know you are looping through the entire list you will use this

	for i := range days {
		fmt.Println(days[i])
	}

	for index, value := range days { // if you want to ignore any one of that you can use _
		fmt.Println("Index:", index, "Value:", value)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 1 {
			goto gotoStatement // jump on to this
		}
		if rougueValue == 5 {
			break
		}
		if rougueValue == 3 { // this is how to use continue in go first increament the value and continue
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++

	}

	// this is a lable
gotoStatement:
	fmt.Println("Jumping at gotoStatement right now ")
}
