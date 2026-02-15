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

// This approach we follow in many recursions; like pick and not pick
// example of real world is like when we try clothes, to wear new onle we need to remove old
// following the same in below approach too
func PowersetRecur2(array []int) [][]int {
	subsets := [][]int{}
	powersetHelper2(array, 0, []int{}, &subsets)
	return subsets
}

func powersetHelper2(array []int, index int, tempArr []int, subsets *[][]int) {
	if index >= len(array) {
		*subsets = append(*subsets, tempArr)
		return
	}

	tempArr = append(tempArr, array[index]) // pick
	powersetHelper2(array, index+1, tempArr, subsets)
	tempArr = tempArr[:len(tempArr)-1] // not pick
	powersetHelper2(array, index+1, tempArr, subsets)
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/25_powerset
func powersetRecurMain() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(PowersetRecur(arr))
	fmt.Println(PowersetRecur2(arr))
}
