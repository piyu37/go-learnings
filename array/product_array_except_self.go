package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	product := 1
	for i := range nums {
		result[i] = product
		product *= nums[i]
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * product
		product *= nums[i]
	}

	return result
}

// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150
func productArrayExceptSelfMain() {
	arr := []int{4, 5, 1, 8, 2}
	fmt.Println(productExceptSelf(arr))
}
