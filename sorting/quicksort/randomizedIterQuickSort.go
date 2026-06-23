package main

import "fmt"

func randomizedIterQuickSort(array []int) []int {
	low := 0

	high := len(array) - 1

	if low >= high {
		return array
	}

	top := -1
	stack := make([]int, high+1)
	top++
	stack[top] = low
	top++
	stack[top] = high

	for top > 0 {
		high = stack[top]
		top--
		low = stack[top]
		top--

		p := partition(array, low, high)

		if low < p-1 {
			top++
			stack[top] = low
			top++
			stack[top] = p - 1
		}

		if high > p+1 {
			top++
			stack[top] = p + 1
			top++
			stack[top] = high
		}
	}

	return array
}

func randomizedIterQuickSortMain() {
	fmt.Println(randomizedIterQuickSort([]int{8}))
}
