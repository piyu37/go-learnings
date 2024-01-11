package main

import "fmt"

func NextGreaterElement(array []int) []int {
	result := make([]int, 0)

	for range array {
		result = append(result, -1)
	}

	stack := make([]int, 0)

	for i := 0; i < len(array); i++ {
		// circularIdx := i % len(array)

		for len(stack) > 0 && array[i] > array[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = array[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}

// https://leetcode.com/problems/next-greater-element-ii/description/
func nextGreaterElementMain() {
	arr := []int{2, 5, -3, -4, 6, 7, 2}

	fmt.Println(NextGreaterElement(arr))
}
