package main

import "fmt"

func proAdd(values ...int) int {

	total := 0
	for _, value := range values {
		total += value

	}
	return total
}

func main() {

	total := proAdd(1, 2, 3)
	fmt.Println("Total: ", total)

}
