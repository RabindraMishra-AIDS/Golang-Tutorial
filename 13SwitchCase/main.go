package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Creating a Dice Game using Switch Case.
func main() {
	fmt.Println("Welcome to the Dice Game using Switch Case")
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1 //Will ensure non zero number
	fmt.Println("Value of Dice is:", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Now You can Start!!")
	case 2:
		fmt.Println("Move 2 Spots")
	case 3:
		fmt.Println("Move 3 Spots")
	case 4:
		fmt.Println("Move 4 spots")
		fallthrough  //After 4 we will also print for 5 but in five we have not written fallthrough so it will stop after printing 5.
	case 5:
		fmt.Println("Move 5 spots")
	case 6:
		fmt.Println("Move 6 spots and play Again!!")
	default:
		fmt.Println("Unexpected error!!")
	}
}
