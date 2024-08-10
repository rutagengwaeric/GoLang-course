package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race condition")

	// Race condition
	wg := &sync.WaitGroup{}
	var score = []int{0}
	var mut = &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One Race")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two Race")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three Race")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
