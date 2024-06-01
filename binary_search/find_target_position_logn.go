package main

import "fmt"

func BinarySearch(array []int, target int) []int {
	leftIdx := alteredBinarySearch(array, target, true)
	rightIdx := alteredBinarySearch(array, target, false)

	return []int{leftIdx, rightIdx}
}

func alteredBinarySearch(array []int, target int, goLeft bool) int {
	idx := -1
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2

		if array[mid] == target {
			idx = mid
			if goLeft {
				if mid > 0 && array[mid-1] == target {
					right = mid - 1
				} else {
					return idx
				}
			} else {
				if mid < len(array)-1 && array[mid+1] == target {
					left = mid + 1
				} else {
					return idx
				}
			}
		} else if target < array[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return idx
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
// WAP to find the range of target value in sorted array
// output: [4, 9]; target = 45
func findTargetPositionLogN() {
	arr := []int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}
	fmt.Println(BinarySearch(arr, 0))
}
