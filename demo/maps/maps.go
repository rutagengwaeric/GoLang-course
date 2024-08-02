package main

import "fmt"

func main() {

	shopingList := make(map[string]int)

	shopingList["eggs"] = 11
	shopingList["milk"] = 1
	shopingList["bread"] += 1
	shopingList["eggs"] += 1

	fmt.Println(shopingList)

	delete(shopingList, "bread")

	fmt.Println("Milk deleted  :", shopingList)
	fmt.Println("need ", shopingList["eggs"], "eggs")

	cereal, found := shopingList["eggs"]

	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Founded : ", cereal)
	}

	totalItems := 0

	for _, value := range shopingList {
		totalItems += value
	}

	fmt.Println("Total Items :", totalItems)

}
