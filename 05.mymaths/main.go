package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
)

func main() {

	// var num1 int = 5
	// var num2 float64 = 4
	// fmt.Println(num1 + int(num2))in this we loose precision

	// random number
	// rand.Seed is depricaiated

	// fmt.Println(rand.Intn(5))

	// random from crypto

	myRandomnum , _ := rand.Int(rand.Reader,big.NewInt(5)) // governed by cryptography randomness is much more accurate
	fmt.Print(myRandomnum)
}
