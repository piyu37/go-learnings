package main

import "fmt"

func MaxSubsetSumNoAdjacent(arr []int) int {
	result := 0

	if len(arr) == 0 {
		return result
	}

	if len(arr) == 1 {
		return arr[0]
	}

	max1 := arr[0]
	max2 := arr[1]

	for i := 2; i < len(arr); i++ {
		max1 = arr[i-2] + arr[i]

		if i-3 >= 0 {
			max2 = arr[i-3] + arr[i]
		}

		arr[i] = max(max1, max2)
	}

	result = max(max1, max2)

	return result
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/13_max_subset_sumno_adjacent
func maxSubsetSumNoAdjacentMain() {
	arr := []int{75, 105, 120, 75, 90, 135}

	fmt.Println(MaxSubsetSumNoAdjacent(arr))
	fmt.Println(MaxSubsetSumNoAdjacentWithoutModifyingInput(arr))
}
