package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please rate pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil { // strconv.ParseFloat: parsing "4\n": invalid syntax \n is a trailing character so need to trim it so use package strings
		fmt.Println(err)
	} else {
		fmt.Println(numRating+1)
	}

}
