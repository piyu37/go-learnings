package main

import "fmt"

func StaircaseTraversal(height int, maxSteps int) int {

	return staircaseHelper(height, maxSteps, map[int]int{0: 1, 1: 1})
}

func staircaseHelper(h, maxSteps int, cache map[int]int) (n int) {
	if h < 0 {
		return 0
	}

	if v, ok := cache[h]; ok {
		return v
	}

	for i := 1; i <= maxSteps; i++ {
		output := staircaseHelper(h-i, maxSteps, cache)

		cache[h-i] = output

		n += output
	}

	return n
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/45_staircase_traversal
func staircaseTraversalMain() {
	fmt.Println(StaircaseTraversal(5, 3))
}
