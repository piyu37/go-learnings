package main

import (
	"fmt"
	"sort"
)

// https://www.geeksforgeeks.org/sort-even-numbers-ascending-order-sort-odd-numbers-descending-order/
func sortOddEvenWhenArrayNotSorted(arr []int) []int {
	countPos := 0
	// Make all odd numbers negative
	for i := 0; i < len(arr); i++ {
		if arr[i]&1 != 0 { // Check for odd
			// Counting only positive odd numbers
			if arr[i] > 0 {
				countPos++
			}
			arr[i] *= -1
		}
	}

	// Sort positive odd and all even numbers
	sort.Ints(arr)

	// Retaining the original array
	for i := 0; i < len(arr); i++ {
		if arr[i]&1 != 0 {
			arr[i] *= -1
		}
	}

	// Sort all the numbers.
	sort.Ints(arr[countPos:])

	return arr
}

// above method variant
// without using extra space and in O(n) time complexity
func sortOddEvenWhenArraySorted(arr []int) []int {
	countIdx := 0
	mulModDivNo := 100

	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i] % mulModDivNo
		if val&1 != 0 {
			arr[countIdx] = (val * mulModDivNo) + arr[countIdx]
			countIdx++
		}
	}

	for i := 0; i < len(arr); i++ {
		val := arr[i] % mulModDivNo
		if val&1 == 0 {
			arr[countIdx] = (val * mulModDivNo) + arr[countIdx]
			countIdx++
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] /= mulModDivNo
	}

	return arr
}

func sortOddEven() {
	arr := []int{1, 2, 3, 5, 4, 7, 10, -3}
	fmt.Println(sortOddEvenWhenArrayNotSorted(arr))

	arr2 := []int{1, 2, 3, 4, 5, 7, 10} 
	fmt.Println(sortOddEvenWhenArraySorted(arr2))
}
