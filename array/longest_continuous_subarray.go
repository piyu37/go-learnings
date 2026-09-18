package main

import "fmt"

func longestSubarrayUsingMonotonicQueue(nums []int, limit int) int {
	minQ := make([]int, 0) // mono increasing
	maxQ := make([]int, 0) // mono decreasing
	longestSubArrayCount := 0

	left, right := 0, 0
	for right < len(nums) {
		// removing from back/last until nums[right] is great than last val; mono increasing
		for len(minQ) != 0 && nums[right] < nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}

		minQ = append(minQ, right)

		// removing from back/last until nums[right] is smaller than last val; mono decreasing
		for len(maxQ) != 0 && nums[right] > nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}

		maxQ = append(maxQ, right)

		for abs(nums[maxQ[0]]-nums[minQ[0]]) > limit {
			if minQ[0] == left {
				minQ = minQ[1:] // removing from start
			}

			if maxQ[0] == left {
				maxQ = maxQ[1:] // removing from start
			}

			left++
		}

		if right-left+1 > longestSubArrayCount {
			longestSubArrayCount = right - left + 1
		}

		right++
	}

	return longestSubArrayCount
}

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/
func longestContinuousSubarray() {
	fmt.Println(longestSubarrayUsingMonotonicQueue([]int{1, 8, 6, 10}, 8))
	fmt.Println(longestSubarrayUsingMonotonicQueue([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarrayUsingMonotonicQueue([]int{1, 5, 6, 7, 8, 10, 6, 5, 6}, 4))
	fmt.Println(longestSubarrayUsingMonotonicQueue([]int{1, 8, 6, 10}, 8))
}
