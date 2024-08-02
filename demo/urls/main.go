package main

import (
	"fmt"
	"net/url"
)

// main is the entry point of the program
func main() {
	// Declare a constant string variable to hold the URL

	const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjghjg"

	// Print the URL
	fmt.Println(myurl)

	// Parse the URL using the url.Parse function
	result, _ := url.Parse(myurl)

	// Print the raw query part of the URL

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// Get the query parameters of the URL as a map
	qureyparams := result.Query()
	fmt.Println(qureyparams)
	// fmt.Println(qureyparams["coursename"])

	// Iterate over the query parameters and print each parameter
	for _, value := range qureyparams {
		// Print each parameter with its value
		fmt.Println("Param is :", value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutss",
		RawPath: "user=hitech",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
