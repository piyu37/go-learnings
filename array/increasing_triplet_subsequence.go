package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first := math.MaxInt32
	second := math.MaxInt32
	for i := range nums {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/increasing-triplet-subsequence/
func increasingTripletSubsequence() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums))
}
