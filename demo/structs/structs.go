package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}
type Bus struct {
	FrontSeat Passenger
}

func main() {

	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		brian = Passenger{Name: "Brian", TicketNumber: 2}
		eric  = Passenger{Name: "Eric", TicketNumber: 3}
	)

	fmt.Println(brian, eric)

	var heidi Passenger

	heidi.Name = "Heidi"
	heidi.Boarded = true
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	brian.Boarded = true

	fmt.Println(casey)

	if casey.Boarded && brian.Boarded {
		fmt.Println(casey.Name, "and", brian.Name, "have a bus")
	}
	heidi.Boarded = true
	fmt.Println(heidi)

	bus := Bus{heidi}
	fmt.Println(bus)

	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

}
