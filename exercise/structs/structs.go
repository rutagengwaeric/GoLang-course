// --Summary:
//
//	Create a program to calculate the area and perimeter
//	of a rectangle.
//
// --Requirements:
// * Create a rectangle structure containing its coordinates
package main

import "fmt"

type Cordinates struct {
	x int
	y int
}

type Rectangle struct {
	a Cordinates // top left
	b Cordinates // bottom right
}

func width(rect Rectangle) int {
	return rect.b.x - rect.a.x
}

func length(rect Rectangle) int {
	return rect.a.y - rect.b.y
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal

func area(rect Rectangle) int {
	return width(rect) * length(rect)
}

func perimeter(rect Rectangle) int {
	return 2 * (width(rect) + length(rect))
}

//  - The functions must use the rectangle structure as the function parameter

//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

func printInfo(rect Rectangle) {
	fmt.Println("Rectangle: ", rect)
	fmt.Println("Area: ", area(rect))
	fmt.Println("Perimeter: ", perimeter(rect))
}

func main() {

	rect := Rectangle{a: Cordinates{0, 7}, b: Cordinates{10, 0}}
	printInfo(rect)
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.a.y *= 2
	rect.b.x *= 2

	//  - Print the new results to the terminal
	printInfo(rect)

}
