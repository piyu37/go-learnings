package main

import "fmt"

func NumberOfBinaryTreeTopologies(n int) int {
	return calculateTopology(n, map[int]int{
		0: 1,
		1: 1,
	})
}

func calculateTopology(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	var count int

	for i := 0; i < n; i++ {
		left := calculateTopology(i, cache)
		right := calculateTopology(n-i-1, cache)

		count += left * right
	}

	cache[n] = count

	return count
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/29_number_of_binary_tree_topologies
func binaryTopology() {
	fmt.Println(NumberOfBinaryTreeTopologies(3))
}
