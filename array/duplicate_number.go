package main

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
