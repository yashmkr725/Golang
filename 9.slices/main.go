package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"Apple","tomato","peach"}
	fmt.Printf("Type of fruis is %T\n", fruits) //Type of fruis is []string
	
	// it automatically expands its size 
	fruits = append(fruits, "Banana") // adding data 
	fmt.Println(fruits)

	//slices
	fruits = append(fruits[:2])
	fmt.Println(fruits)
	fruitlist := append(fruits[1:3])
	fmt.Println(fruitlist)
	
	highScore := make([]int,4)
	// it is the defualt allcation of memory 
	highScore[0] = 73
	highScore[1] = 72
	highScore[2] = 71
	highScore[3] = 70
	// highScore[4] = 70 index out of range [4] with length 4

	highScore = append(highScore, 55,66,21) // this is not error  when we use the method append it reallocate the memory and all values are accomodated entire memory allocation happen again

	fmt.Println(highScore)

	sort.Ints(highScore)//[73 72 71 70 55 66 21]
	fmt.Println(highScore)//[21 55 66 70 71 72 73]
	fmt.Println(sort.IntsAreSorted(highScore)) //true

	// remove values from slice based on index
	var courses = []string{"react","node","rust","python"}
	index := 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}
