package main

import "fmt"

func main() {

	defer fmt.Println("Hello")
	fmt.Println("World")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
