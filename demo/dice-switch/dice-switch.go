package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switchcase Dice Game")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("You got 1")
	case 2:
		fmt.Println("You got 2")
	case 3:
		fmt.Println("You got 3")
	case 4:
		fmt.Println("You got 4")
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	default:
		fmt.Println("You got", diceNumber)
	}

}
