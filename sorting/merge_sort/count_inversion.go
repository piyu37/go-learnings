package main

import "fmt"

func countInversion(array []int) int {
	if len(array) == 0 {
		return 0
	}

	count := 0
	doMergeSort(array, 0, len(array)-1, &count)

	return count
}

func doMergeSort(arr []int, start, end int, count *int) []int {
	if start >= end {
		return []int{arr[start]}
	}

	mid := (end-start)/2 + start
	leftSortedArr := doMergeSort(arr, start, mid, count)
	rightSortedArr := doMergeSort(arr, mid+1, end, count)

	i := 0
	j := 0
	leftSortedArrLen := len(leftSortedArr)
	rightSortedArrLen := len(rightSortedArr)

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
			*count += leftSortedArrLen - i
			j++
		}
	}

	return result
}

// https://www.geeksforgeeks.org/problems/inversion-of-array-1587115620/1
func countInversionMain() {
	arr := []int{2, 4, 1, 3, 5}
	fmt.Println(countInversion(arr))

	arr = []int{5, 3, 2, 4, 1}
	fmt.Println(countInversion(arr))
}
