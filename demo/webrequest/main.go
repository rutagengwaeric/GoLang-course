package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	const url = "https://jsonplaceholder.typicode.com/posts?_limit=3"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of http is %T\n", response)

	databytes, err := ioutil.ReadAll(response.Body)

	content := string(databytes)

	if err != nil {
		panic(err)
	}
	fmt.Println(content)

}
