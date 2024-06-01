package main

import (
	"fmt"
	"math"
)

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	target := math.MaxInt32

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			target = nums[mid]
		}

		// which means left part from is sorted
		if nums[mid] >= nums[left] {
			// checking target is in left or not otherwise right
			if target < nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// checking target is in right or not otherwise left
			if target > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return target
}

func findMin2(nums []int) int {
	// If the list has just one element then return that element.
	if len(nums) == 1 {
		return nums[0]
	}

	// Initializing left and right pointers.
	left := 0
	right := len(nums) - 1

	// If the last element is greater than the first element then there is no rotation.
	// E.g. 1 < 2 < 3 < 4 < 5 < 7. Already sorted array.
	// Hence the smallest element is first element. A[0]
	if nums[right] > nums[0] {
		return nums[0]
	}

	// Binary search way
	for right >= left {
		// Find the mid element
		mid := left + (right-left)/2

		// If the mid element is greater than its next element then mid+1 element is the smallest
		// This point would be the point of change. From higher to lower value.
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		// If the mid element is lesser than its previous element then mid element is the smallest
		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		// If the mid elements value is greater than the 0th element this means
		// the least value is still somewhere to the right as we are still dealing with elements greater than nums[0]
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			// If nums[0] is greater than the mid value then this means the smallest value is somewhere to the left
			right = mid - 1
		}
	}

	return -1
}

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func minRotatedArrSearch() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums))

	nums = []int{5, 1, 2, 3, 4}
	fmt.Println(findMin2(nums))

	nums = []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))

	nums = []int{11, 13, 15, 17}
	fmt.Println(findMin(nums))
}
