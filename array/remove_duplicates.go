package main

import "fmt"

func shiftLeft(arr []int, left, right int) []int {
	if left == right {
		return arr
	}

	element := arr[left]

	// Shift elements to the left
	for i := left; i <= right; i++ {
		arr[i-1] = arr[i]
	}

	// Move the first element to the end
	arr[right] = element

	return arr
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	last := -100000
	count := 0
	output := 0
	right := n - 1
	i := 0
	for i <= right {
		if nums[i] != last {
			count = 1
			last = nums[i]
			output++
			i++
		} else {
			count++
			if count < 3 {
				output++
				i++
			} else {
				nums = shiftLeft(nums, i, right)
				right--
			}
		}
	}

	return output
}

func removeDuplicatesMain() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
