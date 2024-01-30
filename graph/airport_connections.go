package main

import "fmt"

func get_min_no_of_airport_connections(airports []string, routes [][]string, startingAirport string) int {
	routeMap := make(map[string][]string)
	connectionMap := make(map[string][]string)

	for _, route := range routes {
		routeMap[route[0]] = append(routeMap[route[0]], route[1])
	}

	for _, airport := range airports {
		if destinations, ok := routeMap[airport]; ok {
			delete(routeMap, airport)
			for _, d := range destinations {
				connectionDFS(airport, d, routeMap, connectionMap)
			}
		}
	}

	return len(connectionMap)
}

func connectionDFS(startingAirport, destnAirport string, routeMap, connectionMap map[string][]string) {
	if _, ok := connectionMap[destnAirport]; ok {
		if startingAirport != destnAirport {
			connectionMap[startingAirport] = append(connectionMap[startingAirport], destnAirport)
			connectionMap[startingAirport] = append(connectionMap[startingAirport], connectionMap[destnAirport]...)
			delete(connectionMap, destnAirport)
		}
	} else {
		connectionMap[startingAirport] = append(connectionMap[startingAirport], destnAirport)
	}

	if destinations, ok := routeMap[destnAirport]; ok {
		delete(routeMap, destnAirport)
		for _, d := range destinations {
			connectionDFS(startingAirport, d, routeMap, connectionMap)
		}
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/22_airport_connections
func airportConnectionsMain() {
	airports := []string{"BGI", "CDG", "DEL", "DOH", "DSM", "EWR", "EYW", "HND", "ICN",
		"JFK", "LGA", "LHR", "ORD", "SAN", "SFO", "SIN", "TLV", "BUD"}

	routes := [][]string{
		{"DSM", "ORD"},
		{"ORD", "BGI"},
		{"BGI", "LGA"},
		{"SIN", "CDG"},
		{"CDG", "SIN"},
		{"CDG", "BUD"},
		{"DEL", "DOH"},
		{"DEL", "CDG"},
		{"TLV", "DEL"},
		{"EWR", "HND"},
		{"HND", "ICN"},
		{"HND", "JFK"},
		{"ICN", "JFK"},
		{"JFK", "LGA"},
		{"EYW", "LHR"},
		{"LHR", "SFO"},
		{"SFO", "SAN"},
		{"SFO", "DSM"},
		{"SAN", "EYW"},
	}

	startingAirport := "LGA"

	fmt.Println(get_min_no_of_airport_connections(airports, routes, startingAirport))
	fmt.Println(AirportConnections(airports, routes, startingAirport))
}
