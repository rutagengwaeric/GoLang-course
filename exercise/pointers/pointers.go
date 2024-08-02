package main

import "fmt"

//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}
func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(item []Item) {
	for i := 0; i < len(item); i++ {
		deactivate(&item[i].tag)
	}
}

func main() {

	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	//  - Deactivate any one security tag in the array/slice
	//  - Call the checkout() function to deactivate all tags
	//  - Print out the array/slice after each change

	//  - Create at least 4 items, all with active security tags

	shirts := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	shoes := Item{"Shoes", Active}
	hats := Item{"Hats", Active}

	items := []Item{shirts, pants, shoes, hats}

	fmt.Println(items)

	//  - Deactivate any one security tag in the array/slice

	deactivate(&items[0].tag)

	fmt.Println("Item 0 deactivated", items)

	//  - Call the checkout() function to deactivate all tags

	checkout(items)

	//  - Print out the array/slice after each change

	fmt.Println("All items checked out", items)

}
