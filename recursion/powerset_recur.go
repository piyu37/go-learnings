package main

import "fmt"

func PowersetRecur(array []int) [][]int {
	return powersetHelper(array, len(array)-1)
}

func powersetHelper(array []int, index int) [][]int {
	if index < 0 {
		return [][]int{{}}
	}

	subsets := powersetHelper(array, index-1)

	val := array[index]

	for j := range subsets {
		currentSubset := subsets[j]
		newSubset := append([]int{}, currentSubset...)
		newSubset = append(newSubset, val)
		subsets = append(subsets, newSubset)
	}

	return subsets
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/25_powerset
func powersetRecurMain() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(PowersetRecur(arr))
}
