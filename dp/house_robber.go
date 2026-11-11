package main

import "fmt"

func rob(nums []int) int {
	memoization := make([]int, len(nums))

	for i := range memoization {
		memoization[i] = -1
	}

	return robHouse(nums, 0, memoization)

}

func robHouse(nums []int, curr int, memoization []int) int {
	if curr >= len(nums) {
		return 0
	}

	if memoization[curr] >= 0 {
		return memoization[curr]
	}

	robbingCurrent := nums[curr] + robHouse(nums, curr+2, memoization)

	robbingWithoutCurrent := robHouse(nums, curr+1, memoization)

	maxRobbery := max(robbingCurrent, robbingWithoutCurrent)

	memoization[curr] = maxRobbery

	return maxRobbery
}

func rob2(nums []int) int {
	// prev2 = best total up to two houses back
	// prev1 = best total up to the previous house
	prev2, prev1 := 0, 0

	for _, money := range nums {
		// Two choices at this house:
		//   skip it -> keep prev1
		//   rob it  -> prev2 + money (can't rob the adjacent house)
		best := max(prev1, prev2+money)
		prev2, prev1 = prev1, best
	}
	return prev1
}

// https://leetcode.com/problems/house-robber/description/?envType=study-plan-v2&envId=top-interview-150
func houseRobber() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
	fmt.Println(rob2(nums))
}
