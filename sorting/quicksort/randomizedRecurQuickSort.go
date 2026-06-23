package main

import (
	"fmt"
	"math/rand"
)

func partition(array []int, start, end int) int {
	p := start + rand.Intn(end-start+1)
	pVal := array[p]

	array[p], array[end] = array[end], array[p]

	i := start - 1

	for j := start; j < end; j++ {
		if array[j] <= pVal {
			i++

			array[i], array[j] = array[j], array[i]
		}
	}

	i++

	array[i], array[end] = array[end], array[i]

	return i
}

func randomizedQuickSort(array []int, start, end int) []int {
	if start < end {
		p := partition(array, start, end)

		array = randomizedQuickSort(array, start, p-1)
		array = randomizedQuickSort(array, p+1, end)
	}

	return array
}

func randomizedRecurQuickSort() {
	fmt.Println(randomizedQuickSort([]int{8, 5, 2, 9, 5, 6, 3}, 0, 6))
}
