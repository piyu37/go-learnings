package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	monotonicMaxDeque := make([]int, 0, len(nums))

	result := make([]int, 0, len(nums)-k+1)

	for i := range nums {
		// removing first element from monotonicMaxDeque when i is max & it is out of bound from sliding window
		if len(monotonicMaxDeque) != 0 && i == monotonicMaxDeque[0]+k {
			monotonicMaxDeque = monotonicMaxDeque[1:]
		}

		for len(monotonicMaxDeque) != 0 && nums[i] >= nums[monotonicMaxDeque[len(monotonicMaxDeque)-1]] {
			monotonicMaxDeque = monotonicMaxDeque[:len(monotonicMaxDeque)-1]
		}

		monotonicMaxDeque = append(monotonicMaxDeque, i)

		if i+1 >= k {
			result = append(result, nums[monotonicMaxDeque[0]])
		}
	}

	if k > len(nums) {
		result = append(result, nums[monotonicMaxDeque[0]])
	}

	return result
}

// https://leetcode.com/problems/sliding-window-maximum/
func slidingWindowMaximumMain() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}
