package main

import (
	"fmt"
	"math/rand"
)

func partiton(array []int, low, high int) int {
	pivot := rand.Intn(high-low+1) + low

	array[high], array[pivot] = array[pivot], array[high]

	i := low - 1
	pivot = array[high]

	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++

			array[i], array[j] = array[j], array[i]
		}
	}

	i++

	array[high], array[i] = array[i], array[high]

	return i
}

func quickSort(array []int, low, high, k int) int {
	if low <= high {
		p := partiton(array, low, high)

		if p+1 == k {
			return array[p]
		}

		if k < p+1 {
			return quickSort(array, low, p-1, k)
		} else {
			return quickSort(array, p+1, high, k)
		}
	}

	return -1
}

func Quickselect(array []int, k int) int {
	high := len(array) - 1
	low := 0

	return quickSort(array, low, high, k)

}

func quickSelectMain() {
	arr := []int{43, 24, 37}
	fmt.Println(Quickselect(arr, 2))
}
