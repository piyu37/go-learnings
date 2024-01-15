package main

import "fmt"

func MergeSort(array []int) []int {
	arr := applyMergeSort(array, 0, len(array))
	return arr
}

func applyMergeSort(arr []int, start, end int) []int {
	if start == end {
		if start >= len(arr) {
			return []int{}
		}

		return []int{arr[start]}
	}

	mid := (end-start)/2 + start
	leftSortedArr := applyMergeSort(arr, start, mid)
	rightSortedArr := applyMergeSort(arr, mid+1, end)

	i := 0
	j := 0
	leftSortedArrLen := len(leftSortedArr)
	rightSortedArrLen := len(rightSortedArr)
	maxLen := leftSortedArrLen
	if rightSortedArrLen > maxLen {
		maxLen = rightSortedArrLen
	}

	result := make([]int, 0)

	for i < leftSortedArrLen || j < rightSortedArrLen {
		if i >= leftSortedArrLen {
			result = append(result, rightSortedArr[j])
			j++
			continue
		}

		if j >= rightSortedArrLen {
			result = append(result, leftSortedArr[i])
			i++
			continue
		}

		if leftSortedArr[i] < rightSortedArr[j] {
			result = append(result, leftSortedArr[i])
			i++
		} else {
			result = append(result, rightSortedArr[j])
			j++
		}
	}

	return result
}

func mergeSortMain() {
	arr := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(MergeSort(arr))
}
