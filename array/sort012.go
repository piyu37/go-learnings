package main

import "fmt"

func sort012(nums []int) {
	left, mid := 0, 0
	right := len(nums) - 1

	for mid <= right {
		if nums[mid] == 0 {
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		} else if nums[mid] == 2 {
			nums[right], nums[mid] = nums[mid], nums[right]
			right--
		} else {
			mid++
		}
	}
}

// https://leetcode.com/problems/sort-colors/description/
func sort012Main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sort012(nums)
	fmt.Println(nums)
}
