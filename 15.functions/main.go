package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greet()

	result := resultGiver("Fail")
	fmt.Println("Your Result is :", result)

	proResult, proString := proAdder(1, 2, 3, 4, 5, 5, 5, 3, 2)
	fmt.Println(proResult)
	fmt.Println(proString)
}

// means there are values comming of n amout
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi there "
}
func resultGiver(value string) string {
	return "Pass" + value
}
func greet() {
	fmt.Println("Namaste")
}
