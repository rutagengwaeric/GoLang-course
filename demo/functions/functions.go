package main

import "fmt"

func double(x int) int {
	return x * x
}
func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello world")
}

func main() {

	greet()

	result := double(10)
	fmt.Println("The result is", result)

	sum := add(result, 20)
	fmt.Println("The sum is", sum)

}
