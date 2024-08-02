//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string // Book Title

type Name string // Library Member Name

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}
type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // Total owned by the library
	lended int // Total books Lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member Member) {

	for title, audit := range member.books {
		var returnTime string

		if audit.checkIn.IsZero() {
			returnTime = "Not returned"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, audit.checkOut.String(), "Through", returnTime)

	}
}

func main() {

}
