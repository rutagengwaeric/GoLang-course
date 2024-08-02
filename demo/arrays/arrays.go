package main

import "fmt"

type Rooms struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Rooms) {

	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Printf("Room %s is clean\n", room.name)
		} else {
			fmt.Printf("Room %s is dirty\n", room.name)
		}
	}

}

func main() {

	rooms := [...]Rooms{
		{name: "Living Room"},
		{name: "Kitchen"},
		{name: "Bedroom"},
		{name: "Bathroom"},
	}

	checkCleanliness(rooms)
	fmt.Println("Perfoming cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness(rooms)

}
