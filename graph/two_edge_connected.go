package main

import "fmt"

func isTwoEdgeConnected(edges [][]int) bool {
	if len(edges) == 0 {
		return true
	}

	arrivalTimes := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		arrivalTimes[i] = -1
	}

	return getMinimumArrivalTimeOfAncestors(0, -1, 0, arrivalTimes, edges) == -1
}

func getMinimumArrivalTimeOfAncestors(currentVertex, parent, currentTime int, arrivalTimes []int, edges [][]int) int {
	minArrivalTime := currentTime

	arrivalTimes[currentVertex] = currentTime

	for _, destination := range edges[currentVertex] {
		if arrivalTimes[destination] == -1 {
			minArrivalTime = min(minArrivalTime, getMinimumArrivalTimeOfAncestors(destination,
				currentVertex, currentTime+1, arrivalTimes, edges))
		} else if destination != parent {
			minArrivalTime = min(minArrivalTime, arrivalTimes[destination])
		}
	}

	if minArrivalTime == currentTime && parent != -1 {
		return -1
	}

	return minArrivalTime
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/21_two_edge_connected_graph
func twoEdgeConnected() {
	edges := [][]int{
		{5, 1, 2},
		{0, 2},
		{0, 1, 3},
		{2, 4, 5},
		{3, 5},
		{0, 3, 4},
	}

	fmt.Println(isTwoEdgeConnected(edges))
}
