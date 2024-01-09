package main

func IndexEqualsValue(array []int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2

		if mid == array[mid] {
			if mid > 0 && array[mid-1] == mid-1 {
				right = mid - 1
				continue
			} else {
				return mid
			}
		}

		if array[mid] < mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
