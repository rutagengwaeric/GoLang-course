package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {

	fmt.Println("Getting Started with MongoDB and Go! Using the Collection API")

	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8080...")

}
