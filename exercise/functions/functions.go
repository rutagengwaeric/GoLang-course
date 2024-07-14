// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:

package main

import "fmt"

func number(a int) int {
	return a
}

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.

func greet(name string) {
	fmt.Println("Good evening " + name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()

func getMessage() string {
	return "Hello, World!"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer

func add(a int, b int, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func getNumber() int {
	return 10
}

//--Notes:
//* Write a function that returns any two numbers

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

func main() {

	greet("Eric")

	fmt.Println(getMessage())

	sum := add(number(10), 20, 30)
	fmt.Println(sum)

	fmt.Println(getNumber())

}
