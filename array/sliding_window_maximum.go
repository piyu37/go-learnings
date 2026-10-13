package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	dq := make([]int, 0, len(nums))

	result := make([]int, 0, len(nums)-k+1)

	for i := range nums {
		if len(dq) != 0 && i == dq[0]+k {
			dq = dq[1:]
		}

		for len(dq) != 0 {
			if nums[i] < nums[dq[len(dq)-1]] {
				break
			}

			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)

		if i+1 >= k {
			result = append(result, nums[dq[0]])
		}
	}

	if k > len(nums) {
		result = append(result, nums[dq[0]])
	}

	return result
}

// https://leetcode.com/problems/sliding-window-maximum/
func slidingWindowMaximumMain() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}
