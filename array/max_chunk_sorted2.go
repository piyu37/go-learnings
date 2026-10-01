package main

import "fmt"

func maxChunksToSorted2(arr []int) int {
	right := make([]int, len(arr))

	minVal := arr[len(arr)-1]
	right[len(arr)-1] = minVal
	for i := len(arr) - 2; i >= 0; i-- {
		minVal = min(minVal, arr[i])
		right[i] = minVal
	}

	chunks := 1
	maxVal := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		maxVal = max(maxVal, arr[i])

		if maxVal <= right[i+1] {
			chunks++
		}
	}

	return chunks
}

// https://leetcode.com/problems/max-chunks-to-make-sorted-ii/
func maxChunksToSorted2Main() {
	fmt.Println(maxChunksToSorted2([]int{2, 1, 3, 4, 4}))
	fmt.Println(maxChunksToSorted2([]int{1, 0, 1, 3, 2}))
	fmt.Println(maxChunksToSorted2([]int{0, 0, 1, 1, 1}))
}
