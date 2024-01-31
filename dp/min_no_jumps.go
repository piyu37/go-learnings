package main

import "fmt"

func findMinNoOfJumps(arr []int) int {
	jumps := 0
	maxReach := arr[0]
	steps := arr[0]
	for i := 1; i < len(arr); i++ {
		maxReach = max(maxReach, arr[i]+i)
		steps--

		if steps == 0 {
			if i != len(arr)-1 {
				jumps++
			}
			steps = maxReach - i
		}
	}

	return jumps + 1
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/12_min_number_of_jumps
func minNoJumps() {
	arr := []int{1, 1}
	fmt.Println(findMinNoOfJumps(arr))
}
