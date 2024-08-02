package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {

	hello := "Hello"
	world := "World"

	counter := Counter{}
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)

	fmt.Println(hello, world)

	phrase := []string{hello, world}

	replace(&phrase[1], "GO!", &counter)

	fmt.Println(phrase)

}
