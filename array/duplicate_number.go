package main

import "fmt"

// this will only work when all elements in array is less than len(elements)
func findDuplicates(nums []int) []int {
	var duplicates []int

	// Mark elements as visited
	for i := 0; i < len(nums); i++ {
		index := nums[i] % len(nums)
		nums[index] += len(nums)
	}

	// Check for duplicates
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 2*len(nums) {
			duplicates = append(duplicates, i)
		}
	}

	return duplicates
}

// find duplicates from array when numbers in array is lesser than length of array
func duplicateNumber() {
	nums := []int{0, 4, 3, 2, 7, 8, 2, 3, 1}
	duplicates := findDuplicates(nums)
	fmt.Println("The repeating elements are:")
	for _, num := range duplicates {
		fmt.Println(num)
	}
}
