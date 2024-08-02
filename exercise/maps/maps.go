//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//

package main

import "fmt"

//--Requirements:
//* Create a function to print server status displaying:

func printServerStatus(servers map[string]int) {
	fmt.Println("\n---- total name servers :", len(servers))
	//  - number of servers

	stats := make(map[int]int)

	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}
	fmt.Println("Online servers is", stats[Online])
	fmt.Println("Offline servers is", stats[Offline])
	fmt.Println("Maintenance servers is", stats[Maintenance])
	fmt.Println("Retired servers is", stats[Retired])

	//  - number of servers for each status (Online, Offline, Maintenance, Retired)
}

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	serverStatus := make(map[string]int)

	for _, server := range servers {
		//* Set all of the server statuses to `Online` when creating the map
		serverStatus[server] = Online

	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	printServerStatus(serverStatus)

	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired

	//  - change server status of `aiur` to `Offline`

	serverStatus["aiur"] = Offline
	//  - call display server info function
	printServerStatus(serverStatus)

	//  - change server status of all servers to `Maintenance`

	for key, _ := range serverStatus {
		serverStatus[key] = Maintenance
	}
	//  - call display server info function

	printServerStatus(serverStatus)

}
