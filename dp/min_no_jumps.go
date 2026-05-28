package main

import "fmt"

func findMinNoOfJumps(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	jumps := 0
	maxReach := arr[0]
	steps := arr[0]
	for i := 1; i < len(arr); i++ {
		maxReach = max(maxReach, arr[i]+i)
		steps--

		if steps == 0 {
			if i != len(arr)-1 {
				jumps++
			}
			steps = maxReach - i
		}
	}

	return jumps + 1
}

func findMinNoOfJumps2(nums []int) int {
	right := nums[0]

	if len(nums) < 2 {
		return 0
	}

	i := 0
	jumpCount := 1
	totalCount := right
	for i <= right {
		if i == len(nums)-1 {
			return jumpCount
		}

		right = max(right, i+nums[i])
		i++
		if i > totalCount {
			jumpCount++
			totalCount = right
		}
	}

	return jumpCount
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/12_min_number_of_jumps
func minNoJumps() {
	arr := []int{1, 1}
	fmt.Println(findMinNoOfJumps(arr))

	arr = []int{1, 1}
	fmt.Println(findMinNoOfJumps2(arr))
}
