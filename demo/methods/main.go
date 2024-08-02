package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) getStatus() {
	fmt.Println("Your Status is: ", u.Status)

}

func main() {

	user := User{"Eric", "CfGQH@example.com", 20, true}
	fmt.Printf("Hello %v Your email is %v and you are %v years old\n", user.Name, user.Email, user.Age)

	fmt.Printf("All details %+v\n", user)

	user.getStatus()
}
