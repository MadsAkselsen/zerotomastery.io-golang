//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

func printServerStatus(servers map[string]int) {
	fmt.Println("Total Servers:", len(servers))

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
			fmt.Println("Unknown status")
		}
	}
	fmt.Println("Online:", stats[Online])
	fmt.Println("Offline:", stats[Offline])
	fmt.Println("Maintenance:", stats[Maintenance])
	fmt.Println("Retired:", stats[Retired])
}

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serversMap := make(map[string]int)

	for _, server := range servers {
		serversMap[server] = Online
	}

	fmt.Println("==========")
	printServerStatus(serversMap)

	serversMap["darkstar"] = Retired
	serversMap["aiur"] = Offline

	fmt.Println("==========")
	printServerStatus(serversMap)

	for i := 0; i < len(servers); i++ {
		serverName := servers[i]
		serversMap[serverName] = Maintenance
	}

	fmt.Println("==========")
	printServerStatus(serversMap)
}
