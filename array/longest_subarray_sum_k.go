package main

import "fmt"

func longestSubArraySumK(arr []int, k int) int {
	sumMap := map[int]int{}
	maxLen := 0
	sum := 0

	for i := range arr {
		sum += arr[i]

		if sum == k {
			len := i + 1
			if len > maxLen {
				maxLen = len
			}
		}

		if v, ok := sumMap[sum-k]; ok {
			len := i - v
			if len > maxLen {
				maxLen = len
			}
		}

		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}

	return maxLen
}

// https://takeuforward.org/data-structure/longest-subarray-with-given-sum-k/
// code where we are getting both negative & positive values
func longestSubArraySumKMain() {
	arr := []int{1, 0, -1}
	fmt.Println(longestSubArraySumK(arr, -1))
}
