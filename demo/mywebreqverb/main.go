package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to go lang course web request verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformFormRequest()
}

func PerformGetRequest() {

	const url = "http://localhost:8000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println("response status is ", response.StatusCode)
	fmt.Println("response Lenght is ", response.ContentLength)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	var ResponseString strings.Builder
	byteCount, _ := ResponseString.Write(content)

	fmt.Println("lenght is ", byteCount)
	fmt.Println(ResponseString.String())

	if err != nil {
		panic(err)
	}

}

func PerformPostJsonRequest() {

	const url = "http://localhost:8000/post"
	const jsonData = `
	{
		"course":"golang",
		"name":"eric"
	}`

	response, err := http.Post(url, "application/json", strings.NewReader(jsonData))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}

func PerformFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Rutagengwa")
	data.Add("lastname", "Eric")
	data.Add("email", "eric@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
