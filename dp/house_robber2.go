package main

import "fmt"

// Houses are in a circle, so the first and last house can't both be robbed.
// That leaves two cases, each of which is plain House Robber on a line:
//   1. skip the last house  -> rob2(nums[:n-1])
//   2. skip the first house -> rob2(nums[1:])
// The answer is the better of the two.
// Time: O(n), Space: O(1)
func robCircle(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0] // only one house, nothing to conflict with
	}
	return max(rob2(nums[:n-1]), rob2(nums[1:]))
}

// https://leetcode.com/problems/house-robber-ii/description/
func houseRobber2() {
	fmt.Println(robCircle([]int{2, 3, 2}))    // 3
	fmt.Println(robCircle([]int{1, 2, 3, 1})) // 4
	fmt.Println(robCircle([]int{1, 2, 3}))    // 3
}
