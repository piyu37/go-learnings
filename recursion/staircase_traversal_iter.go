package main

import "fmt"

func StaircaseTraversalIter(height int, maxSteps int) int {
	result := make([]int, height+1)

	result[0], result[1] = 1, 1

	sum := 1

	for i := 2; i <= height; i++ {
		sum = sum + result[i-1]
		if i-maxSteps-1 >= 0 {
			sum -= result[i-maxSteps-1]
		}

		result[i] = sum
	}

	return result[height]
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/45_staircase_traversal
func staircaseTraversalIterMain() {
	fmt.Println(StaircaseTraversalIter(5, 3))
}
