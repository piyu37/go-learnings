package main

import "fmt"

func findMostPowerfulAlliance(arr []int) int {
	msis := make([]int, len(arr))
	maxSum := 0

	copy(msis, arr)

	for i := 2; i < len(arr); i++ {
		for j := 0; j < i-1; j++ {
			if msis[i] < arr[i]+msis[j] {
				msis[i] = arr[i] + msis[j]
				if maxSum < arr[i]+msis[j] {
					maxSum = arr[i] + msis[j]
				}
			}
		}
	}

	return maxSum
}

// variant/similar of max sum increasing subsequence
// https://github.com/lee-hen/Algoexpert/tree/master/hard/10_max_sumIncreasing_subsequence
func mostPowerfulAlliance() {
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println(findMostPowerfulAlliance(arr))
}
