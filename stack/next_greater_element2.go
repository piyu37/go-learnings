package main

import "fmt"

func nextGreaterElements2(nums []int) []int {
	result := make([]int, len(nums))

	for i := range result {
		result[i] = -1
	}

	n := len(nums)
	stack := make([]int, 0)
	for i := (2 * n) - 1; i >= 0; i-- {
		idx := i % n

		for len(stack) > 0 && nums[idx] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if result[idx] == -1 && len(stack) > 0 {
			result[idx] = nums[stack[len(stack)-1]]
		}

		stack = append(stack, idx)
	}

	return result
}

// https://leetcode.com/problems/next-greater-element-ii/description/
func nextGreaterElement2() {
	arr := []int{3, 8, 4, 1, 2}
	fmt.Println(nextGreaterElements2(arr))

	arr = []int{1, 2, 3, 4, 3}
	fmt.Println(nextGreaterElements2(arr))
}
