package main

import "fmt"

func NextGreaterElement(array []int) []int {
	result := make([]int, 0)

	for range array {
		result = append(result, -1)
	}

	stack := make([]int, 0)

	for i := range array {
		for len(stack) > 0 && array[i] > array[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = array[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}

func nextGreaterElementOriginal(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)

	for i, v := range nums1 {
		nums1Map[v] = i
	}

	result := make([]int, len(nums1))
	for i := range result {
		result[i] = -1
	}

	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if v, ok := nums1Map[nums2[i]]; ok && len(stack) > 0 {
			result[v] = stack[len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	return result
}

// https://leetcode.com/problems/next-greater-element-i/description/
func nextGreaterElementMain() {
	arr := []int{2, 5, -3, -4, 6, 7, 2}

	fmt.Println(NextGreaterElement(arr))

	nums1 := []int{1, 3, 4, 2}
	nums2 := []int{4, 1, 2}

	fmt.Println(nextGreaterElementOriginal(nums1, nums2))
}
