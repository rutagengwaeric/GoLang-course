package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")

	weblists := []string{
		"https://roadready.rw",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range weblists {
		go getStatusCode(web)
		wg.Add(1)

	}

	wg.Wait()

	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Ooops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("Status code for %s is %d\n", endpoint, res.StatusCode)
	}

}
