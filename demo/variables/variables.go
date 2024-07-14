package main

import "fmt"

func main() {
	var myName = "Eric"
	fmt.Println("My name is ", myName)

	var name string = "Jack"
	fmt.Println("My name is ", name)

	a := 10
	fmt.Println("a is ", a)

	var sum int
	fmt.Println("sum is ", sum)

	part1, other := 20, 3
	fmt.Println("part1 is ", part1, "other is ", other)

	part2, other := 30, 4
	fmt.Println("part2 is ", part2, "other is ", other)

	sum = 44
	fmt.Println("sum is ", sum)

	var names string = "Rutagengwa"
	fmt.Println("My name is", names)

	userName := "Admin"
	fmt.Println(`Your Username is ` + userName)

	part, other := 10, 100
	fmt.Println("part1 is ", part, "other is", other)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName is", lessonName, "lessonType is", lessonType)

}
