// --Summary:
//
//	Create a program that can store a shopping list and print
//	out information about the list.
//
// --Requirements:
//   - Using an array, create a shopping list with enough room
//     for 4 products
//   - Products must include the price and the name
//   - Insert 3 products into the array
//   - Print to the terminal:
//   - The last item on the list
//   - The total number of items
//   - The total cost of the items
//   - Add a fourth product to the list and print out the
//     information again
package main

import "fmt"

type Product struct {
	price int
	name  string
}

func printInfo(products [4]Product) {
	var totalCost, totalNumber int
	for i := 0; i < len(products); i++ {

		totalCost += products[i].price

		if products[i].name != "" {
			totalNumber += 1
		}
		// product := products[i]
		// fmt.Println(product)
	}

	fmt.Println("Last Product: ", products[totalNumber-1])
	fmt.Println("Total Cost: ", totalCost)
	fmt.Println("Total Number: ", totalNumber)

}

func main() {

	products := [4]Product{
		{price: 10, name: "Product 1"},
		{price: 20, name: "Product 2"},
		{price: 30, name: "Product 3"},
	}

	printInfo(products)

	products[3] = Product{price: 40, name: "Product 4"}

	printInfo(products)

}
