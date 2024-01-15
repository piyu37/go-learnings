package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(array []int, start, end int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	p := start + r.Intn(end-start+1)

	array[p], array[end] = array[end], array[p]

	i := start - 1

	p = array[end]

	for j := start; j < end; j++ {
		if array[j] <= p {
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
