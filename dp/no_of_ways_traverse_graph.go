package main

import "fmt"

func findNoOfWaysToTraverseGraph(w, h int) int {
	currArr := make([]int, w)
	for i := range currArr {
		currArr[i] = 1
	}

	for i := 1; i < h; i++ {
		prevArr := make([]int, len(currArr))
		copy(prevArr, currArr)
		prevBlockVal := 1

		for j := 1; j < w; j++ {
			currArr[j] = prevArr[j] + prevBlockVal
			prevBlockVal = currArr[j]
		}
	}

	return currArr[w-1]
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/40_number_of_ways_to_traverse_graph
func noOfWaysTraverseGraph() {
	fmt.Println(findNoOfWaysToTraverseGraph(2, 3))
	fmt.Println(NumberOfWaysToTraverseGraph2(4, 3))
}
