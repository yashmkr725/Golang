package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case")
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of Dice is", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is one you can Open")
	case 2:
		fmt.Println("Dice value is two you can Move")
	case 3:
		fmt.Println("Dice value is three you can Move")
		fallthrough // it is used when if there is a usecase when the value hit 3 then i need also case 4 op so it execute both case 3 and 4
		// op
		// Value of Dice is 3
		// Dice value is three you can Move
		// Dice value is four you can Move
	case 4:
		fmt.Println("Dice value is four you can Move")
	case 5:
		fmt.Println("Dice value is five you can Move")
	case 6:
		fmt.Println("Dice value is six you can Move")
	default:
		fmt.Println("Illegal")
	}
	// Switch case Done !
}
