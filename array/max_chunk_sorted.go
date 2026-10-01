package main

import "fmt"

func maxChunksToSorted(arr []int) int {
	maxVal := -1
	chunks := 0

	for i := range arr {
		maxVal = max(maxVal, arr[i])

		if i == maxVal {
			chunks++
		}
	}

	return chunks
}

// https://leetcode.com/problems/max-chunks-to-make-sorted/description/
func maxChunksToSortedMain() {
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}
