package main

import "fmt"

func buildMaxHeap(arr []int) []int {
	parent := (len(arr) - 2) / 2

	for i := parent; i >= 0; i-- {
		arr = shiftDown(arr, i, len(arr)-1)
	}

	return arr
}

func shiftDown(arr []int, start, end int) []int {
	childIdx1 := (2 * start) + 1

	for childIdx1 <= end {
		childIdx2 := (2 * start) + 2

		var maxIdx int

		if childIdx2 > end {
			maxIdx = childIdx1
		} else {
			maxIdx = childIdx1

			if arr[childIdx1] < arr[childIdx2] {
				maxIdx = childIdx2
			}
		}

		if arr[maxIdx] > arr[start] {
			arr[maxIdx], arr[start] = arr[start], arr[maxIdx]
		}

		start = maxIdx
		childIdx1 = start*2 + 1
	}

	return arr
}

func HeapSort(array []int) []int {
	array = buildMaxHeap(array)

	for i := 0; i < len(array); i++ {
		length := len(array) - i - 1
		array[0], array[length] = array[length], array[0]
		array = shiftDown(array, 0, length-1)
	}

	return array
}

func heapSortMain() {
	arr := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(HeapSort(arr))
}
