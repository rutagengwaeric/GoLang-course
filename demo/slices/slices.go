package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println("----", title, "----")
	for i := 0; i < len(slice); i++ {
		iterator := slice[i]
		fmt.Println(iterator)
	}

}

func main() {

	route := []string{"Delhi", "Mumbai", "Chennai", "Kolkata", "Pune", "Jaipur"}
	printSlice("Route 1", route)
	route = append(route, "Hyderabad")

	printSlice("Route 2", route)

	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	fmt.Println("---- Remaining Routes----")

	fmt.Println(route[2:])

	var numbers []int

	for k := 0; k < 100; k++ {
		numbers = append(numbers, k)
	}

	for i := 0; i < len(numbers); i++ {

		output := numbers[i]
		fmt.Println(output)
	}

}
