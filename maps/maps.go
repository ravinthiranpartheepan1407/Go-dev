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

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	fmt.Println(servers)
	mapped := make(map[string]int)
	mapped["darkstar, aiur"] = Online
	println(mapped)

	// fruits := []string{"Apple", "Orange", "Mango"}
	maps := make(map[string]int)
	maps["Apple"] = 1
	maps["ornage"] = 2
	fmt.Println(maps)

	server := Online

	for _, ch := range mapped {
		server += ch
	}

	fmt.Println(server)

}
