package main

import "fmt"

func recursiveQuickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		recursiveQuickSort(arr, low, p-1)
		recursiveQuickSort(arr, p+1, high)
	}
}

func recursiveQuickSortMain() {
	arr := []int{4, 3, 5, 1, 2, 3, 1, 3}

	arrLen := len(arr)

	recursiveQuickSort(arr, 0, arrLen-1)

	fmt.Println(arr)

}
