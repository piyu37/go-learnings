package main

import "fmt"

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	for i := range nums {
		numMap[nums[i]] = true
	}

	longestCount := 0
	for k, v := range numMap {
		if !v {
			continue
		}

		numMap[k] = false

		rCount, lCount := 0, 0
		kFront := k
		kBack := k

		for numMap[kFront+1] {
			numMap[kFront+1] = false
			rCount++
			kFront++
		}

		for numMap[kBack-1] {
			numMap[kBack-1] = false
			lCount++
			kBack--
		}

		if 1+lCount+rCount > longestCount {
			longestCount = lCount + rCount + 1
		}
	}

	return longestCount
}

// https://leetcode.com/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-interview-150
func longestConsecutiveSequence() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
