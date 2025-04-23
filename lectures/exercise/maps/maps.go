//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServers(serverList map[string]int) {

	online := 0
	offline := 0
	maintenance := 0
	retired := 0
	total := 0

	for _, status := range serverList {
		if status == Online {
			online += 1
		} else if status == Offline {
			offline += 1
		} else if status == Maintenance {
			maintenance += 1
		} else {
			retired += 1
		}

		total += 1

	}

	fmt.Printf("Total servers %v; online servers %v; offline servers %v; servers in maintenance %v; retired servers %v \n", total, online, offline, maintenance, retired)

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serversList := make(map[string]int)
	for _, name := range servers {
		serversList[name] = Online
	}

	displayServers(serversList)

	serversList["darkstar"] = Retired
	serversList["aiur"] = Offline
	displayServers(serversList)

	for name, _ := range serversList {
		serversList[name] = Maintenance
	}

	displayServers(serversList)
}
