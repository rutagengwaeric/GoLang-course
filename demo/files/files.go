package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to go lang files")

	content := "This needs to go in a files"

	file, err := os.Create("./myfile.txt")

	checkNillErr(err)

	length, err := io.WriteString(file, content)

	checkNillErr(err)

	fmt.Println("Length is: ", length)

	defer file.Close()
	readFile("./myfile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNillErr(err)

	fmt.Println(string(databyte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
